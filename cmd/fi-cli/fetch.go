package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"
)

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

func fetch(year string) {
	url, ok := urlForYear[year]
	if !ok {
		fmt.Printf("Missing data for year %s!\n", year)
		return
	}

	fmt.Printf("Fetching from %s\n", url)
	zipFile := fmt.Sprintf("doc-%s.zip", year)
	kmlFile := fmt.Sprintf("doc-%s.kml", year)
	// download da fuckin' zip
	if err := downloadFile(zipFile, url); err != nil {
		panic(err)
	}

	// extract da fuckin' zip
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name != "doc.kml" {
			continue
		}

		outFile, err :=
			os.OpenFile(kmlFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return
		}

		rc, err := f.Open()
		if err != nil {
			return
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
	}
	syscall.Unlink(zipFile)

	fmt.Println("SUCCESS")
}
