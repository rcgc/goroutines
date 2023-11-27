package main

import (
	"fmt"
	"time"
)

/*
	Usually you'll see a default clause used in conjuction
	with a for-select loop.

	This allows a goroutine to make progress on work while
	waiting for another goroutine to report a result.

	Finally, the following statement will block forever:
	select{}
*/

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(done)
	}()

	workCounter := 0
	loop:
	for {
		select {
		case <- done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1*time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop. \n", workCounter)
}