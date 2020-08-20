/**
@author: Vikas K
**/
package core

import (
	"log"
	"testing"
)

var xmlRaw = `
< type="PUBLIC">
<name>vikas</name>
	<address>
		<street>2585 Park Rd 6026</street>
		<city>Johnson City</city>
		<state>TX</state>c
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

func TestGetXmlMap(t *testing.T) {
	receivingXmlMap, err := GetXmlMap(xmlRaw, "->")
	if err != nil {
		log.Println("Err: ", err)
		return
	} else if receivingXmlMap == nil {
		log.Println("Received nil xml map")
		return
	}
	for key, value := range receivingXmlMap {
		log.Println(key, " > ", value)
	}
}
