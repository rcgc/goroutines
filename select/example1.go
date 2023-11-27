package main

import "fmt"

/*
	The select statement is the glue that binds channels together
*/
func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
	select {
	case <- c1:
		// Do something
		fmt.Print("Channel 1")
	case <- c2:
		// Do something
		fmt.Print("Channel 2")
	case c3<- struct{}{}:
		// Do something
		fmt.Print("Channel 3")
	}
}