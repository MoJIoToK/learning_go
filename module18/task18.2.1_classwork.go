package main

import (
	"fmt"
	"sync"
)

const COUNT_ROUTINES = 5
const PRINT_TIMES = 10

func printNumber(routine int) {
	for i := 0; i < PRINT_TIMES; i++ {
		fmt.Println(routine)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(COUNT_ROUTINES)
	for routine := 1; routine <= COUNT_ROUTINES; routine++ {
		go func(routine int) {
			printNumber(routine)
			wg.Done()
		}(routine)
	}
	wg.Wait()
}
