// This example uses function ParseGlob https://golang.org/pkg/text/template/#ParseGlob
// As the name says this takes a glob of templates.
// So instead of ParseFiles which takes in list of all the files as input parameter, we will give all the files together in ParseGlob.

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// This will return a pointer to a template. https://golang.org/pkg/text/template/#ParseGlob
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// This ExecuteTemplate allows us to chose which template we want to execute.
	err = tpl.ExecuteTemplate(os.Stdout, "b.stupid", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// This ExecuteTemplate allows us to chose which template we want to execute.
	err = tpl.ExecuteTemplate(os.Stdout, "c.stupid", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Since now the tpl contains multiple template, using Execute will print the first one that was added to the template. You can change the order and verify.
	// Execute takes in a Writer Interface and since os.Stdout implemens the Writer interface, it can be used here as an input.
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
