package main

import (
	"encoding/xml"
)

type KmlRoot struct {
	XMLName  xml.Name    `xml:"kml"`
	Document KmlDocument `xml:"Document"`
}

type KmlDocument struct {
	XMLName     xml.Name      `xml:"Document"`
	Name        string        `xml:"name"`
	Description string        `xml:"description"`
	Styles      []KmlStyle    `xml:"Style"`
	StyleMaps   []KmlStyleMap `xml:"StyleMap"`
	Folders     []KmlFolder   `xml:"Folder"`
}

type KmlStyleMap struct {
	XMLName xml.Name `xml:"StyleMap"`
	Id      string   `xml:"id,attr"`
}

type KmlStyle struct {
	XMLName xml.Name `xml:"Style"`
	Id      string   `xml:"id,attr"`
}

type KmlFolder struct {
	XMLName    xml.Name       `xml:"Folder"`
	Name       string         `xml:"name"`
	Placemarks []KmlPlacemark `xml:"Placemark"`
}

type KmlPlacemark struct {
	XMLName     xml.Name `xml:"Placemark"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
}
