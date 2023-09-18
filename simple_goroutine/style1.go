package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
	// continue doing other things
}

/* 
	There's a risk that this doesn't print anything
 	because of race condition. So, that's why
	sleep has been added (race condition still exists)
*/
func main() {
	go sayHello()
	// time.Sleep(2 * time.Second)
}