package main

import (
	"fmt"
	"sync"
)

/*
	1) We define a type Button that contains a condition, Clicked.

	2) Here we define a convenience function that will allow us to register functions to
	handle signals from a condition. Each handler is run on its own goroutine, and
	subscribe will not exit until that goroutine is confirmed to be running.

	3) Here we set a handler for when the mouse button is raised. It in turn calls Broad
	cast on the Clicked Cond to let all handlers know that the mouse button has
	been clicked (a more robust implementation would first check that it had been
	depressed).

	4) Here we create a WaitGroup. This is done only to ensure our program doesn’t exit
	before our writes to stdout occur.

	5) Here we register a handler that simulates maximizing the button’s window when
	the button is clicked.

	6) Here we register a handler that simulates displaying a dialog box when the mouse
	is clicked.

	7) Next, we simulate a user raising the mouse button from having clicked the appli‐
	cation’s button.
*/

func main() {
	type Button struct { // 1)
		Clicked *sync.Cond
	}

	button := Button{ Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()){ // 2)
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // 3)
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() { // 4)
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	
	subscribe(button.Clicked, func() { // 5)
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() { // 6)
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() // 7)
	clickRegistered.Wait()
}