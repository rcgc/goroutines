package main

import (
	"fmt"
	"sync"
)

/*
	Closures close around the lexical scope they are created in,
	thereby capturing variables.

	goroutines execute within the same address space they
	were created in, and so our program prints out the word “welcome.”
*/

func main(){
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func(){
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}