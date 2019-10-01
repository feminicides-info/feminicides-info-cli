package main

// import "github.com/davecgh/go-spew/spew"

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type ConvertOptions = struct {
	InputKml   string
	OutputJson string
}

var indexOfMonth = map[string]time.Month{
	"janvier":   time.January,
	"fevrier":   time.February,
	"mars":      time.March,
	"avril":     time.April,
	"mai":       time.May,
	"juin":      time.June,
	"juillet":   time.July,
	"aout":      time.August,
	"septembre": time.September,
	"octobre":   time.October,
	"novembre":  time.November,
	"decembre":  time.December,
}

type JsonFeminicides struct {
	Feminicides []JsonWoman `json:"feminicides"`
}

type JsonWoman struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

/* Detect folder containing the most placemarks */
func getDataFolder(kml KmlRoot) *KmlFolder {
	var currentFolder *KmlFolder
	var maxFolder *KmlFolder

	for i := 0; i < len(kml.Document.Folders); i++ {
		currentFolder = &kml.Document.Folders[i]
		if maxFolder == nil {
			maxFolder = currentFolder
		}

		// switch if max found
		if len(currentFolder.Placemarks) > len(maxFolder.Placemarks) {
			maxFolder = currentFolder
		}
	}

	// No folders found !
	return maxFolder
}

func extractName(name string) string {
	// remove NBSP
	fixedName := strings.Replace(name, "\u00a0", " ", -1)
	r, _ := regexp.Compile("^[O\\d]+\\s*-\\s*(\\S+)")
	match := r.FindStringSubmatch(fixedName)
	if len(match) > 0 {
		return match[1]
	}
	fmt.Fprintf(os.Stderr,
		"WARNING: extractName: Unable to extract name from « %s »\n",
		fixedName)
	return "(anonymous)"
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func removeAccents(input string) string {
	output := make([]byte, len(input))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, e := t.Transform(output, []byte(input), true)
	if e != nil {
		panic(e)
	}

	return string(output)
}

func extractDate(desc string, year int) time.Time {
	r, _ := regexp.Compile("((lundi|mardi|mercredi|jeudi|vendredi|samedi|dimanche|week-end du)\\W+)?(\\d+/)?(\\d+)(er)?\\W+(janvier|fevrier|mars|avril|mai|juin|juillet|aout|septembre|octobre|novembre|decembre)")
	match := r.FindStringSubmatch(strings.ToLower(desc))

	month := indexOfMonth[match[6]]
	day, _ := strconv.Atoi(match[4])

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func extractYear(description string) int {
	r, _ := regexp.Compile("^(\\d+)\\s+")
	match := r.FindStringSubmatch(description)
	year, _ := strconv.Atoi(match[1])

	return year
}

func stripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))

	str, _, _ = transform.String(t, str)
	return str
}

func convert(options ConvertOptions) error {
	var murders JsonFeminicides
	var kml KmlRoot
	var err error
	var rawXml []byte

	if options.InputKml == "-" {
		rawXml, err = ioutil.ReadAll(os.Stdin)
	} else {
		rawXml, err = ioutil.ReadFile(options.InputKml)
	}
	cleanRawXml := []byte(stripCtlAndExtFromUnicode(string(rawXml)))

	fmt.Fprintf(os.Stderr, "Converting KML %s into JSON %s...\n", options.InputKml, options.OutputJson)

	if err != nil {
		return errors.New("convert: Unable to read input file")
	}

	xml.Unmarshal(cleanRawXml, &kml)

	dataFolder := getDataFolder(kml)
	if dataFolder == nil {
		return errors.New("convert: Unable to detect data folder")
	}

	year := extractYear(kml.Document.Name)
	if year < 1 {
		return errors.New("convert: Unable to extract document year")
	}

	for i := 0; i < len(dataFolder.Placemarks); i++ {
		name := dataFolder.Placemarks[i].Name
		desc := dataFolder.Placemarks[i].Description

		woman := JsonWoman{
			Name: extractName(name),
			Date: extractDate(removeAccents(desc), year),
		}
		// fmt.Printf("%+v\n", woman)
		murders.Feminicides = append(murders.Feminicides, woman)
	}
	rawJson, err := json.MarshalIndent(murders, "", "  ")

	var fh *os.File
	if options.OutputJson == "-" {
		fh = os.Stdout
	} else {
		fh, err = os.OpenFile(options.OutputJson, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	}
	if _, err = fh.Write(rawJson); err != nil {
		fh.Close() // ignore error; Write error takes precedence
		return err
	}
	fh.Close()
	return err
}
