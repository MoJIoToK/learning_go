package main

import (
	"fmt"
	"sync"
)

// STEP_1 is counter increment step.
const STEP_1 int = 1

// ITERATION_AMOUNT is an amount of goroutines.
const ITERATION_AMOUNT int = 1000

func main() {
	var counter int = 0

	//Create an object which is using as Locker.
	var mu = sync.RWMutex{}

	//Create conditional variable.
	var c = sync.NewCond(&mu)

	// Increment is a variable to which an anonymous function is assigned.
	//This function is incremented counter with a given step. Goroutines are synchronized using by conditional variable.
	increment := func() {
		//Blocks all goroutines but the one, where the increment happens.
		//In fact, an object of type mutex is locked.
		c.L.Lock()
		counter += STEP_1
		//fmt.Println(counter)

		//Unblocks all goroutines.
		c.L.Unlock()
	}
	for i := 1; i <= ITERATION_AMOUNT; i++ {
		go increment()
	}

	//This goroutines checks which counter is equal to ITERATION_AMOUNT.
	go func() {
		defer mu.RUnlock()
		for {
			//Block mutex for read.
			mu.RLock()
			if counter == ITERATION_AMOUNT {
				//Signal sends a signal to the main goroutine.
				c.Signal()
				break
			}
			//Unblock mutex.
			mu.RUnlock()
		}
	}()

	c.L.Lock()

	//Main goroutine is waiting signal from conditional variable.
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
