package main

import (
	"fmt"
	"os"
)

func main() {
	var xmlFile, xsdFile string
	if len(os.Args) > 1 {
		xmlFile = os.Args[1]
		xsdFile = os.Args[2]
	} else {
		fmt.Println("No file specified.")
		fmt.Println("")
		fmt.Println("usage: " + os.Args[0] + " filename.xml filename.xsd")
		fmt.Println("Validate an xml file against an xsd schema")
		return
	}
	xsdValidate(fileContents(xmlFile), fileContents(xsdFile))
}
