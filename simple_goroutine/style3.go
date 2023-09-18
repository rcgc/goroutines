package main

import (
	"fmt"
)

/*
	Alternatively, you can assign the function
	to a variable and call the anonymous function
	like this:

	(Sleep has been added to avoid empty output,
	not a good practice but can be used for this simple
	program and race condition still exists)
*/

func main() {
	sayHello := func() {
		fmt.Println("Hello")
	}

	go sayHello()
	// time.Sleep(time.Second * 2)
	// continue doing other things
}