package rld

import "encoding/xml"

type LeoXML struct {
	XMLName     xml.Name `xml:"xml"`
	Sectionlist struct {
		Sectionsort string `xml:"sectionsort,attr"`
		Section     []struct {
			SctTitle string `xml:"sctTitle,attr"`
			Entry    []struct {
				Side []struct {
					Lang  string `xml:"lang,attr"`
					Words struct {
						Text string   `xml:",chardata"`
						Word []string `xml:"word"`
					} `xml:"words"`
				} `xml:"side"`
			} `xml:"entry"`
		} `xml:"section"`
	} `xml:"sectionlist"`
	Similar struct {
		// Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Side   []struct {
			Text string   `xml:",chardata"`
			Lang string   `xml:"lang,attr"`
			Word []string `xml:"word"`
		} `xml:"side"`
	} `xml:"similar"`
	Ffsynlist struct {
		// Text string `xml:",chardata"`
		Word string `xml:"word,attr"`
		Side []struct {
			// Text string `xml:",chardata"`
			Lang string   `xml:"lang,attr"`
			Word []string `xml:"word"`
		} `xml:"side"`
	} `xml:"ffsynlist"`
}
