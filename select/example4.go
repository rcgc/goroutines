package main

import (
	"fmt"
	"time"
)

/*
	1) This case statement will never become unblocked
	because we're reading from a nil channel
*/
func main() {
	var c <-chan int
	select {
	case <-c: // 1)
	case <-time.After(1 * time.Second): // how to use time outs
		fmt.Println("Timed out.")
	}
}