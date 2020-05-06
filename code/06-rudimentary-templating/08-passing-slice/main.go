package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	weekdays := []string{"Monday", "Tuedsay", "Wednesday", "Thursday", "Friday"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", weekdays)
	if err != nil {
		log.Fatalln(err)
	}
}
