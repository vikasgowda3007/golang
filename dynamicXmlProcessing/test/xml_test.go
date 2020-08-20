package test

import (
	"../core"
	"github.com/vikasgowda3007/golang/dynamicXmlProcessing/core"
	"log"
	"testing"
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

func TestXML(t *testing.T) {
	finalMap, err := core.PreFinal(xmlRaw, ">")
	if err != nil {
		log.Println("Err: ", err)
	} else if finalMap == nil {
		return
	}
	for key, value := range finalMap {
		log.Println(key, " > ", value)
	}
}
