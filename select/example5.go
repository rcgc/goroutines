package main

import (
	"fmt"
	"time"
)

/*
	It ran the default statement almost instantaneously.

	This allows you to exit a select block wihtout blocking.
*/

func main() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}