package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FetchOptions = struct {
	Year      string
	OutputKml string
}

var urlForYear = map[string]string{
	"2016": "https://www.google.com/maps/d/u/0/kml?mid=1vikwsH56LM9t5eBu3VOAhzLEqIk",
	"2017": "https://www.google.com/maps/d/u/0/kml?mid=15ZxvoBUgO4ttolUFSoU5UmuRe98",
	"2018": "https://www.google.com/maps/d/u/0/kml?mid=19gV1RSgQ5LNG51BeE-WiV7G8m1MocfXZ",
	"2019": "https://www.google.com/maps/d/u/1/kml?mid=1Y9bFj8Cjfl3rKwuyDBB5-LNkdKKAjtq9",
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func fetch(options FetchOptions) error {
	url, ok := urlForYear[options.Year]
	fmt.Fprintf(os.Stderr, "Fetching KML for year %s\n", options.Year)
	if !ok {
		return errors.New(fmt.Sprintf("Missing data for year %s!\n", options.Year))
	}

	fmt.Fprintf(os.Stderr, "Fetching from %s\n", url)
	zipHandler, err := ioutil.TempFile(os.TempDir(), "fi-cli-*.zip")
	if err != nil {
		log.Fatal(err)
	}
	zipFile := zipHandler.Name()
	defer os.Remove(zipFile)

	kmlFile := options.OutputKml
	if err := downloadFile(zipFile, url); err != nil {
		panic(err)
	}

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name != "doc.kml" {
			continue
		}

		var outFile *os.File
		if kmlFile == "-" {
			outFile = os.Stdout
		} else {
			outFile, err = os.OpenFile(kmlFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
	}
	fmt.Fprintf(os.Stderr, "SUCCESS\n")
	return nil
}
