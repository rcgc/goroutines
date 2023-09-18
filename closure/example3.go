package main

import (
	"fmt"
	"sync"
)

/*
	1) Here we declare a parameter, just like
	any other function. We shadow the original
	salutation variable to make what’s happening
	more apparent.

	2) Here we pass in the current iteration’s
	variable to the closure. A copy of the string
	struct is made, thereby ensuring that when
	the goroutine is run, we refer to the
	proper string.
*/
func main() {
	var wg sync.WaitGroup
	for _, salutation := range [] string {"hello", "greetings", "good day"}{
		wg.Add(1)
		go func(salutation string){ // 1)
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // 2)
	}
	wg.Wait()
}