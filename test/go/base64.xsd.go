// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// MyType1 ...
type MyType1 []byte

// MyType2 ...
type MyType2 struct {
	XMLName    xml.Name `xml:"myType2"`
	LengthAttr int      `xml:"length,attr,omitempty"`
	Value      []byte   `xml:",chardata"`
}

// MyType3 ...
type MyType3 struct {
	XMLName    xml.Name  `xml:"myType3"`
	LengthAttr int       `xml:"length,attr,omitempty"`
	Value      time.Time `xml:",chardata"`
}

// MyType4 ...
type MyType4 struct {
	XMLName   xml.Name `xml:"myType4"`
	Title     string   `xml:"title"`
	Blob      []byte   `xml:"blob"`
	Timestamp string   `xml:"timestamp"`
}

// MyType5 ...
type MyType5 time.Time

// MyType6 ...
type MyType6 struct {
	XMLName xml.Name `xml:"myType6"`
	Value   string   `xml:",chardata"`
}

// TopLevel ...
type TopLevel struct {
	XMLName xml.Name  `xml:"topLevel"`
	MyType1 []byte    `xml:"myType1"`
	MyType2 *MyType2  `xml:"myType2"`
	MyType3 *MyType3  `xml:"myType3"`
	MyType4 *MyType4  `xml:"myType4"`
	MyType5 time.Time `xml:"myType5"`
	MyType6 *MyType6  `xml:"myType6"`
}
