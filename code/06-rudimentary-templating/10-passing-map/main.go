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

	weekdays := map[string]string{
		"monday":    "måndag",
		"tuedsay":   "tisdag",
		"wednesday": "onsdag",
		"thursday":  "torsdag",
		"friday":    "fredag",
	}

	//Go maps do not maintain the insertion order so output probably will not be in same order defined.
	err := tpl.Execute(os.Stdout, weekdays)
	if err != nil {
		log.Fatalln(err)
	}
}
