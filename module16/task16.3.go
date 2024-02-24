package main

import (
	"fmt"
	"time"
)

func main() {
	var result int
	go intervalSums(&result, 0, 50000)
	go intervalSums(&result, 50000, 100000)
	time.Sleep(time.Second) // give goroutines time to finish

	fmt.Println(result)

	otherResult := 0

	for i := 0; i < 100000; i++ {
		otherResult += i
	}
	fmt.Println(otherResult)
}

func intervalSums(destination *int, start, end int) {
	for i := start; i < end; i++ {
		result := *destination
		result += i
		*destination = result
	}
}
