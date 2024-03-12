package main

import (
	"fmt"
	"time"
)

/*
	https://gobyexample.com/select

	Go’s select lets you wait on multiple
	channel operations. Combining goroutines
	and channels with select is a powerful
	feature of Go.

	For our example we’ll select across two channels.

	1) Each channel will receive a value after
	some amount of time, to simulate e.g. blocking RPC
	operations executing in concurrent goroutines.

	2) We’ll use select to await both of these values
	simultaneously, printing each one as it arrives.
*/

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one" // 1)
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two" // 1)
    }()

    for i := 0; i < 2; i++ {
        select { // 2)
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}