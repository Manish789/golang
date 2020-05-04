package main

import "fmt"

func main() {

	// below defer statement is executed, and as part of the execution places fmt.Println("Bye Gopher") on a stack to be executed prior to the function returning.
	defer fmt.Println("Bye Gopher")

	// Below line is executed immediately
	fmt.Println("Hi Gopher")

	// As this the end of the function, fmt.Println("Bye Gopher") is now invoked.
}
