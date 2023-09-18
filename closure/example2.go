package main

import (
	"fmt"
	"sync"
)

/*
	The Go runtime is observant enough to know
	that a reference to the salutation variable
	is still being held, and therefore will transfer
	the memory to the heap so that the goroutines can
	continue to access it.
*/
func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"}{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}