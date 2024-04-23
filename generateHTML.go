package main

import (
	"os"
	"text/template"
)

func generateHTML(data ClinicalDocument) string {
	tmpl, err := template.ParseFiles("./template/template.html")
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("./htmlFiles/result.html")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	err = tmpl.Execute(newFile, data)
	if err != nil {
		panic(err)
	}

	return newFile.Name()
}
