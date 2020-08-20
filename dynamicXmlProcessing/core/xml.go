package core

import (
	"encoding/xml"
	"errors"
	"strings"
)

type xmlTestMap struct {
	Unparsed unparsedTagsMap `xml:",any"`
}

// unparsedTagMap contains the tag information
type unparsedTagMap struct {
	XMLName     xml.Name
	FullContent string `xml:",innerxml"` // for debug purpose, allow to see what's inside some tags
}

// unparsedTagsMap store tags not handled by Unmarshal in a map, it should be labelled with `xml",any"`
type unparsedTagsMap map[string]string

var (
	finalMap  = make(map[string]string)
	separator string
)

func PreFinal(xmlStr string, sep string) (map[string]string, error) {
	var xmlStructMapMain xmlTestMap
	separator = sep
	err := xml.Unmarshal([]byte("<d>"+xmlStr+"</d>"), &xmlStructMapMain)
	if err != nil {
		return nil, err
	}
	if len(xmlStructMapMain.Unparsed) > 1 {
		return nil, errors.New("wrong XML input")
	}
	mainTag := ""
	for key := range xmlStructMapMain.Unparsed {
		mainTag = key
		break
	}
	if mainTag == "" {
		return nil, errors.New("wrong XML input : invalid tag name")
	}
	xmlStructMapMain.Unparsed = nil
	err = xml.Unmarshal([]byte(xmlStr), &xmlStructMapMain)
	if err != nil {
		return nil, err
	}
	err = makeFinalMap(xmlStructMapMain.Unparsed, mainTag)
	if err != nil {
		return nil, err
	}
	return finalMap, nil
}

func makeFinalMap(unparsed unparsedTagsMap, mainTag string) error {
	for key, value := range unparsed {
		if len(value) > 0 && key != "" {
			if strings.Contains(value, "<") {
				var xmlStructMapRec xmlTestMap
				err := xml.Unmarshal([]byte("<"+key+">"+value+"</"+key+">"), &xmlStructMapRec)
				if err != nil {
					return err
				}
				errRec := makeFinalMap(xmlStructMapRec.Unparsed, mainTag+separator+key)
				if errRec != nil {
					return errRec
				}
			} else {
				if mainTag != "" {
					finalMap[mainTag+separator+key] = value
				}
			}
		}
	}
	return nil
}

func (u *unparsedTagsMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *u == nil {
		*u = unparsedTagsMap{}
	}
	e := unparsedTagMap{}
	err := d.DecodeElement(&e, &start)
	if err != nil {
		return err
	}
	//if _, ok := (*u)[e.XMLName.Local]; ok {
	//	return fmt.Errorf("unparsedTagsMap: UnmarshalXML: Tag %s:  multiple entries with the same name", e.XMLName.Local)
	//}
	(*u)[e.XMLName.Local] = e.FullContent
	return nil
}
