package main

import "fmt"

/*
	Lexical confinement involves using
	lexical scope to expose only the correct data
	and concurrency primitives for multiple
	concurrent processes to use.
	It makes it impossible to do the wrong thing.


	1) Here we instantiate the channel within
	the lexical scope of the chanOwner function.
	This limits the scope of the write aspect
	of the results channel to the closure defined
	below it. In other words, it connes the write
	aspect of this channel to prevent other goroutines
	from writing to it.

	2) Here we receive the read aspect of the channel
	and we’re able to pass it into the consumer,
	which can do nothing but read from it. Once again
	this confines the main goroutine to a read-only view
	of the channel.

	3) Here we receive a read-only copy of an int channel.
	By declaring that the only usage we require is read access,
	we confine usage of the channel within the consume function
	to only reads.
*/

func main(){
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // 1)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // 3)
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	
	results := chanOwner() // 2)
	consumer(results)
}