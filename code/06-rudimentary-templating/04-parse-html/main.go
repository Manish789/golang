package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// This will return a pointer to a template. https://golang.org/pkg/text/template/#ParseFiles
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute takes in a Writer Interface and since os.Stdout implemens the Writer interface, it can be used here as an input.
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
