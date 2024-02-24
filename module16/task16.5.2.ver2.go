package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	val := 0
	n = 5

	mu := sync.Mutex{}
	start := time.Now()

	go func() {
		for {
			mu.Lock()
			time.Sleep(time.Second * 1 / 5) // Имитируем некие вычисления
			fmt.Println(val)
			mu.Unlock()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			val++
		}

	}()

	time.Sleep(time.Second * time.Duration(n))
	fmt.Println(val, time.Since(start))

}
