package main

import (
	"fmt"
	"time"
)

var (
	concurrent    = 5
	semaphoreChan = make(chan struct{}, concurrent)
)

func doWork(item int) {
	semaphoreChan <- struct{}{}
	go func() {
		defer func() {
			<-semaphoreChan
		}()
		fmt.Println(item)
		time.Sleep(1 * time.Second)
	}()
}

func main() {
	for i := 0; i < 10000; i++ {
		doWork(i)
	}
}
