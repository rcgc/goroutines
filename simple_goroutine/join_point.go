package main

import (
	"fmt"
	"sync"
)

/*
	Join points are what guarantee our
	program’s correctness and remove
	the race condition.

	This example will deterministically
	block the main goroutine until the goroutine
	hosting the sayHello function terminates.
*/
func main() {
	var wg sync.WaitGroup

	sayHello := func(){
		defer wg.Done()
		fmt.Println("Hello")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait() // This is the join point
}