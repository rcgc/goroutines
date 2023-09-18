package main

import (
	"fmt"
)

/*
	Doing so but using an anonymous function.
	Again, sleep has been added because of
	race condition and there's a risk that
	goroutine prints nothing (not the best
	practice but can be used for this simple case
	and race condition still exists)
*/
func main() {
	go func() {
		fmt.Println("Hello")
	}()
	
	// time.Sleep( time.Second * 2)
}