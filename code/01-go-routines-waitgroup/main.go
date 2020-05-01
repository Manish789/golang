package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello bar", i)
	}

	wg.Done()
}
