package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {

	xmlFolderPath := "./xmlFiles"

	files, err := ioutil.ReadDir(xmlFolderPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xml") {
			xmlFilePath := filepath.Join(xmlFolderPath, file.Name())

			data := readXML(xmlFilePath)

			tempFile := generateHTML(data)
			generatePDF(tempFile, makePdfFilePath(file.Name()))
		}
	}

}

func makePdfFilePath(fileName string) string {
	lastIndex := strings.LastIndex(fileName, ".xml")
	pdfFile := fileName[:lastIndex] + ".pdf" + fileName[lastIndex+4:]

	return "./pdfFiles/" + pdfFile
}
