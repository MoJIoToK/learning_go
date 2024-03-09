package main

import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup // общепринятая практика называть переменную этого примитва как wg

	n := time.Now()
	results := make([]uint, len(computations))
	for i := range computations {
		wg.Add(1) // всегда снаружи рутины
		go func(i int) {
			results[i] = fib(computations[i])
			wg.Done() // всегда внутри рутины
		}(i)
	}

	// заметьте, мы убрали time.Sleep(2*time.Second)

	wg.Wait() // блокируем рутину main и ждём завершения остальных рутин

	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("processed in", time.Since(n))
}
