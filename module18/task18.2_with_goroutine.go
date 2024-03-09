package main

import (
	"fmt"
	"time"
)

// fib — это функция расчитывающая N-ное число последовательноси Фибоначчи
func fib(number uint) uint {
	if number == 0 || number == 1 {
		return number
	}

	return fib(number-2) + fib(number-1)
}

func main() {
	computations := []uint{
		34,
		5,
		12,
		25,
		30,
		42,
		3,
	}

	n := time.Now()
	results := make([]uint, len(computations))
	for i := range computations {
		go func(i int) {
			results[i] = fib(computations[i])
		}(i)
	}

	time.Sleep(2 * time.Second)

	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("processed in", time.Since(n))
}
