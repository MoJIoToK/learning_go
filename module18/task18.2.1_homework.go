package main

import (
	"fmt"
	"sync"
)

const (
	COUNT_ROUTINES = 5
	PRINT_TIMES    = 10
)

// PrintNumber prints serial number of go the goroutine in which it is called.
func printNumber(routine int) {
	for i := 0; i < PRINT_TIMES; i++ {
		fmt.Println(routine)
	}
}

func main() {

	//Create WaitGroup.
	var wg sync.WaitGroup

	//Add 5 goroutine in wait group. Inner atomic count is equal COUNT_ROUTINES.
	wg.Add(COUNT_ROUTINES)
	for routine := 1; routine <= COUNT_ROUTINES; routine++ {
		//Create goroutine
		go func(routine int) {
			printNumber(routine)
			//This method sends signal that the one of the goroutine finished. Inner atomic count decrement.
			wg.Done()
		}(routine)
	}

	//Wait all goroutine in group.
	wg.Wait()
}
