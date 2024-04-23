package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func readXML(filePath string) ClinicalDocument {
	xmlFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var clinicalDocument ClinicalDocument

	xml.Unmarshal(byteValue, &clinicalDocument)

	// exportJson(clinicalDocument)

	return clinicalDocument

}

func exportJson(clinicalDocument ClinicalDocument) {
	file, err := os.Create("console.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(clinicalDocument)
	if err != nil {
		panic(err)
	}
}
