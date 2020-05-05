package main

import (
	"log"
	"os"
	"text/template"
)

// tpl is a pointer to template from the template package
var tpl *template.Template

// init() function sets off a piece of code to run before any other part of your package.
// init() is always called, regardless if there's main or not, so if we import a package that has an init function, it will be executed.
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	// now we are passing data into the template. "Tuesday" in this case.
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Tuesday")
	if err != nil {
		log.Fatalln(err)
	}
}
