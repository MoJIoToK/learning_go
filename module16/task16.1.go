package main

import (
	"fmt"
)

func main() {
	var result int
	go intervalSum(&result, 0, 50)
	intervalSum(&result, 50, 101)

	fmt.Println(result)

	result = 0

	for i := 0; i < 101; i++ {
		result += i
	}
	fmt.Println(result)
}

func intervalSum(destination *int, start, end int) {
	for i := start; i < end; i++ {
		result := *destination
		result += i
		*destination = result
	}
}
