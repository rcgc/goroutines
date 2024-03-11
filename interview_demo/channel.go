package main

import (
	"fmt"
)

type Demo struct {
	value int
}

func worker(d Demo) {
	c := make(chan Demo)
	fmt.Println("I'm worker")

	<-c
}

func main() {
	d := Demo { value: 1 }
	go worker(d)

	//time.Sleep(3 * time.Second)
	fmt.Println("I'm main")
}