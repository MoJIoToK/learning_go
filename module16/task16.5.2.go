package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var n int
	val := 0

	fmt.Println("Enter the duration in second: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	mu := sync.Mutex{}
	start := time.Now()

	go func() {
		for {
			go func() {
				for {
					mu.Lock()
					time.Sleep(time.Second * 1 / 5)
					fmt.Println(val)
					mu.Unlock()
				}
			}()
			time.Sleep(time.Second)
			val++
		}
	}()

	time.Sleep(time.Second * time.Duration(n))
	fmt.Println(val, time.Since(start))

}
