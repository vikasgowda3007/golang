package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"strings"
)

var xmlRaw = `
<campground type="PUBLIC">
<name>vikas</name>
	<address>
		<street>2585 Park Rd 6026</street>
		<city>Johnson City</city>
		<state>TX</state>
		<zip>78636</zip>
	</address>
	<amenities>
		<amenity>
			<distance>Within Facility</distance>
			<name>Biking</name>
		</amenity>
		<amenity>
			<distance>Within Facility</distance>
			<name>Kayaking</name>
		</amenity>
	</amenities>
</campground>`

type xmlTestMap struct {
	Unparsed UnparsedTagsMap `xml:",any"`
}

// UnparsedTagMap contains the tag information
type UnparsedTagMap struct {
	XMLName     xml.Name
	FullContent string `xml:",innerxml"` // for debug purpose, allow to see what's inside some tags
}

// UnparsedTagsMap store tags not handled by Unmarshal in a map, it should be labelled with `xml",any"`
type UnparsedTagsMap map[string]string

var finalMap = make(map[string]string)
var mainTag = ""

func main() {
	err := preFinal(xmlRaw)
	if err != nil {
		log.Println(err)
	}
	for key, value := range finalMap {
		fmt.Println(key, " > ", value)
	}
}

func preFinal(xmlStr string) error {
	var xmlStructMapMain xmlTestMap
	err := xml.Unmarshal([]byte("<d>"+xmlStr+"</d>"), &xmlStructMapMain)
	if err != nil {
		return err
	}
	if len(xmlStructMapMain.Unparsed) > 1 {
		return errors.New("wrong XML input")
	}
	for key := range xmlStructMapMain.Unparsed {
		mainTag = key
		break
	}

	xmlStructMapMain.Unparsed = nil
	err = xml.Unmarshal([]byte(xmlStr), &xmlStructMapMain)
	if err != nil {
		return err
	}
	err = makeFinalMap(xmlStructMapMain.Unparsed, mainTag)
	if err != nil {
		return err
	}
	return nil
}

func makeFinalMap(unparsed UnparsedTagsMap, mainTag string) error {
	for key, value := range unparsed {
		if len(value) > 0 && key != "" {
			if strings.Contains(value, "<") {
				var xmlStructMapRec xmlTestMap
				mainTag = mainTag + ">" + key
				err := xml.Unmarshal([]byte("<"+key+">"+value+"</"+key+">"), &xmlStructMapRec)
				if err != nil {
					return err
				}
				errRec := makeFinalMap(xmlStructMapRec.Unparsed, mainTag)
				if errRec != nil {
					return errRec
				}
			} else {
				if mainTag != "" {
					finalMap[mainTag+">"+key] = value
				}
			}
		}
	}
	return nil
}

func (u *UnparsedTagsMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *u == nil {
		*u = UnparsedTagsMap{}
	}
	e := UnparsedTagMap{}
	err := d.DecodeElement(&e, &start)
	if err != nil {
		return err
	}
	//if _, ok := (*u)[e.XMLName.Local]; ok {
	//	return fmt.Errorf("UnparsedTagsMap: UnmarshalXML: Tag %s:  multiple entries with the same name", e.XMLName.Local)
	//}
	(*u)[e.XMLName.Local] = e.FullContent
	return nil
}
