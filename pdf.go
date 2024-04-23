package main

import (
	"bytes"
	"io/ioutil"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func generatePDF(tempFile string, pdfFile string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	htmlContent, err := ioutil.ReadFile(tempFile)
	if err != nil {
		panic(err)
	}

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlContent)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(pdfFile)
	if err != nil {
		panic(err)
	}

}
