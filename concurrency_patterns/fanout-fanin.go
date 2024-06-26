package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
	1) Here we take in our standard done channel
	to allow our goroutines to be torn down,
	and then a variadic slice of interface{}
	channels to fan-in.

	2) On this line we create a sync.WaitGroup
	so that we can wait until all channels have
	been drained.

	3) Here we create a function, multiplex, which,
	when passed a channel, will read from the channel,
	and pass the value read onto the multiplexedStream
	channel.

	4) This line increments the sync.WaitGroup by the
	number of channels we’re multiplexing.

	5) Here we create a goroutine to wait for all the
	channels we’re multiplexing to be drained so that
	we can close the multiplexedStream channel.
*/
func main() {
	repeatFn := func(done <-chan interface{}, fn func() interface{}, ) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func( done <-chan interface{}, valueStream <-chan interface{}, num int, ) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for integer := range intStream {
				integer -= 1
				prime := true
				for divisor := integer - 1; divisor > 1; divisor-- {
					if integer%divisor == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}, ) <-chan interface{} { // 1)
		var wg sync.WaitGroup // 2)
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) { // 3)
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		// Select from all the channels
		wg.Add(len(channels)) // 4)
		for _, c := range channels {
			go multiplex(c)
		}

		// Wait for all the reads to complete
		go func() { // 5)
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	rand := func() interface{} { return rand.Intn(50000000) }

	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("CPUs available: %d\n", numFinders)
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")

	// Here starts the fanOut part
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	// Here ends the fanOut part

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}