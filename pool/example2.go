package main

import (
	"fmt"
	"sync"
)

/*
	So why use a pool and not just instantiate objects as you go? Go has a garbage
	collector, so the instantiated objects will be automatically cleaned up.

	1) Notice that we are storing the address of the slice of bytes.
	2) And here we are asserting the type is a pointer to a slice of bytes.

	Garbage Collector: it's a system that specifically identifies and cleans up
	dynamic memory allocations.
*/
func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{}{
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem // 1)
		},
	}

	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024*1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte) // 2)
			defer calcPool.Put(mem)
			// Assume something interesting, but quick is being done with
			// this memory.
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}