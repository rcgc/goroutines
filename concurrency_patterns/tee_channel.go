package main

import "fmt"

/*
	https://ornlu-is.github.io/go_tee_channel_pattern/

	1) We’re going to use one select statement so
	that writes to out1 and out2 don’t block each other.
	To ensure both are written to, we’ll perform two
	iterations of the select statement: one for each
	outbound channel.

	2) We will want to use local versions of out1 and out2,
	so we shadow these variables.

	3) Once we’ve written to a channel, we set its shadowed
	copy to nil so that further writes will block and the
	other channel may continue.
*/

func numberStream() <-chan float64 {
	ch := make(chan float64)
	numberStrings := []float64{1., 2., 3., 4., 5., 6., 7., 8., 9., 10.}

	go func() {
		for _, numberString := range numberStrings {
			ch <- numberString
		}

		close(ch)
		return
	}()

	return ch
}

func teeChannel(c <-chan float64) (<-chan float64, <-chan float64) {
	tee1 := make(chan float64)
	tee2 := make(chan float64)

	go func() {
		defer func() {
			close(tee1)
			close(tee2)
		}()

		for val := range c {
			for i := 0; i < 2; i++ { // 1)
				var tee1, tee2 = tee1, tee2 // 2)
				select {
				case tee1 <- val:
					tee1 = nil // 3)
				case tee2 <- val:
					tee2 = nil // 3)
				}
			}
		}

		return
	}()

	return tee1, tee2
}

func main() {
	done := make(chan struct{})
	defer close(done)

	dataStream := numberStream()

	teedStream1, teedStream2 := teeChannel(dataStream)

	for val1 := range teedStream1 {
		fmt.Printf("tee1: %f\n", val1)
		fmt.Printf("tee2: %f\n", <-teedStream2)
	}
}