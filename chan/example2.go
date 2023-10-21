package main

import "fmt"

/*
	1) Here we receive both a string, salutation, and a boolean, ok.
	ok = true, which indicates that channel remains open
*/

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	salutation, ok := <-stringStream // 1)
	fmt.Printf("(%v): %v", ok, salutation)
}