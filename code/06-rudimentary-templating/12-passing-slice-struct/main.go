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

	hanuman := quote{
		Name:  "Hanuman",
		Motto: "Calm, before the struggles and after the storm",
	}

	billGates := quote{
		Name:  "Bill Gates",
		Motto: "Success is a lousy teacher. It seduces smart people into thinking they can't lose",
	}

	quotes := []quote{hanuman, billGates}

	//Go maps do not maintain the insertion order so output probably will not be in same order defined.
	err := tpl.Execute(os.Stdout, quotes)
	if err != nil {
		log.Fatalln(err)
	}
}
