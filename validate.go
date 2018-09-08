package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unsafe"

	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"
)

func fileContents(inFile string) []byte {
	// Read entire file contents into memory, ioutil.ReadFile() closes file after reading.
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Println("Error reading file:", err)
		return contents
	}
	return contents
}

func xsdValidate(xmlContents []byte, xsdContents []byte) {
	xsdSchema, err := xsd.ParseSchema(xsdContents)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc := golibxml.ParseDoc(string(xmlContents))
	if doc == nil {
		// TODO capture and display error
		fmt.Println("Error parsing document")
		return
	}
	defer doc.Free()

	// golibxml._Ctype_xmlDocPtr can't be cast to xsd.DocPtr, even though they are both
	// essentially _Ctype_xmlDocPtr.  Using unsafe gets around this.
	if err := xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("XML Valid as per XSD")
}
