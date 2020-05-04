package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// This will return a pointer to a template. https://golang.org/pkg/text/template/#ParseFiles
	tpl, err := template.ParseFiles("a.stupid")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Later we decide to use the same pointer to template (tpl) from above, we can use the method, https://golang.org/pkg/text/template/#Template.ParseFiles , of the type template to add more templates into tpl.
	tpl, err = tpl.ParseFiles("b.stupid", "c.stupid")
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
