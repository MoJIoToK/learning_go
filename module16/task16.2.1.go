package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var count int

	fmt.Println("Enter gorutines count")
	_, err := fmt.Scan(&count)
	checkerr(err)

	wg := sync.WaitGroup{}

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(j int) {
			for {
				fmt.Println("This goroutine with id - ", i)
				time.Sleep(time.Second)
				defer wg.Done()
			}
		}(i)
	}
	wg.Wait()
}

func checkerr(err error) {
	if err != nil {
		fmt.Println("FATAL: ", err.Error())
		os.Exit(1)
	}
}
