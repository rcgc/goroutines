package main

import (
	"fmt"
)

type Demo struct {
	value int
}

func worker(c chan Demo) {
	fmt.Println("I'm worker")

	c<-Demo{}

	close(c)
}

func main() {
	// d := Demo { value: 1 }
	c := make(chan Demo)
	go worker(c)
	
	//time.Sleep(3 * time.Second)
	fmt.Println("I'm main")

	<-c
}