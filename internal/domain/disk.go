package domain

import "encoding/xml"

type Disk struct {
	XMLName xml.Name `xml:"disk,omitempty"`
	Type    string   `xml:"type,attr,omitempty"`
	Device  string   `xml:"device,attr,omitempty"`

	Driver Driver
	Source Source
	Target Target
}

type Driver struct {
	XMLName xml.Name `xml:"driver,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
	Type    string   `xml:"type,attr,omitempty"`
	Cache   string   `xml:"cache,attr,omitempty"`
}

type Source struct {
	XMLName xml.Name `xml:"source,omitempty"`
	File    string   `xml:"file,attr,omitempty"`
}

type Target struct {
	XMLName xml.Name `xml:"target,omitempty"`
	Dev     string   `xml:"dev,attr,omitempty"`
	Bus     string   `xml:"bus,attr,omitempty"`
}
