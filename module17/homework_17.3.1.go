package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// STEP is counter increment step.
const STEP int64 = 100

// END_COUNTER_VALUE is final value of counter.
const END_COUNTER_VALUE int64 = 1000

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup

	// Increment is a variable to which an anonymous function is assigned.
	//This function is incremented counter with a given step. Goroutines are synchronized using by atomicInt.
	increment := func() {
		defer wg.Done()
		atomic.AddInt64(&counter, STEP)
		//fmt.Println(counter)
	}

	var iterationCount = int(END_COUNTER_VALUE / STEP)
	for i := 1; i <= iterationCount; i++ {
		wg.Add(1)
		go increment()

	}

	// Wait signal.
	wg.Wait()

	// Print result.
	fmt.Println("Result counter -", counter)
}
