package main

import (
	"fmt"
	"sync"
	"time"
)

var mu = sync.Mutex{}

func main() {
	var result int
	go intervalsSum(&result, 0, 5)
	go intervalsSum(&result, 5, 10)
	time.Sleep(time.Second) // give goroutines time to finish

	fmt.Println(result)

	otherResult := 0

	for i := 0; i < 10; i++ {
		otherResult += i
	}
	fmt.Println(otherResult)
}

func intervalsSum(destination *int, start, end int) {
	for i := start; i < end; i++ {
		mu.Lock()
		result := *destination
		result += i
		*destination = result
		fmt.Println(start, end, i, result)
		mu.Unlock()
	}
}
