package main

import (
	"fmt"
	"sync"
)

/*
	because sync.Once only counts the number
	of times Do is called, not how many times
	unique functions passed into Do are called.
	In this way, copies of sync.Once are tightly
	coupled to the functions they are intended
	to be called with; once again we see how
	usage of the types within the sync package work
	best within a tight scope.

	It is recommended to formalize this coupling
	by wrapping any usage of sync.Once in a small
	lexical block:

	either a small function, or by wrapping both in a type.
*/
func main() {
	var count int

	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once

	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}