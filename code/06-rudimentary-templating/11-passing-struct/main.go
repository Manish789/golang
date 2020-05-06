package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type quote struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	billGates := quote{
		Name:  "Bill Gates",
		Motto: "Success is a lousy teacher. It seduces smart people into thinking they can't lose",
	}

	//Go maps do not maintain the insertion order so output probably will not be in same order defined.
	err := tpl.Execute(os.Stdout, billGates)
	if err != nil {
		log.Fatalln(err)
	}
}
