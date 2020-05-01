package main

import (
	"fmt"
)

type person struct {
	Name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.Name, "speaks")
}

func saySomething(h human) {
	h.speak()
}

func main() {

	manish := person{Name: "Manish"}

	// This below call will not work. Uncomment and try it out.
	// saySomething(manish)

	// This call below will work
	saySomething(&manish)
}
