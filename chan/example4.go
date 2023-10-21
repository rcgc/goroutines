package main

import "fmt"

/*
	1) Here we ensure that the channel is closed before we exit the goroutine. This is a
	very common pattern.

	2) Here we range over intStream.

	* The range keyword—used in conjunction with the for statement—supports channels as
	arguments, and will automatically break the loop when a channel is closed.
*/
func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream) // 1)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream { // 2)
		fmt.Printf("%v", integer)
	}
}