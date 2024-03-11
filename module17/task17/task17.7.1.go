package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"task17/counter"
)

// Worker is method for filling the channel. If the counter is greater than or equal max value, recording/sending
// to channel stops.
func worker(ctx context.Context, cancel context.CancelFunc, c *counter.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		//Method for sending message to channel.
		c.Add(1, ctx, cancel)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var wgChan sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	amountOfThreads := 0
	maxValue := 0

	//Enter the count of goroutines.
	fmt.Println("Укажите количество горутин: ")
	_, err := fmt.Scanln(&amountOfThreads)
	if err != nil {
		log.Fatalln("Неверное значение")
	}
	if amountOfThreads < 1 {
		log.Fatalln("Количество горутин не может быть меньше 1")
	}

	//Enter the max value.
	fmt.Println("Укажите максимальное значение счётчика: ")
	_, err = fmt.Scanln(&maxValue)
	if err != nil {
		log.Fatalln("Неверное значение")
	}
	if maxValue < 1 {
		log.Fatalln("Максимальное значение счтчика не может быть меньше 1")
	}

	//Create object - Counter.
	c := counter.NewCounter(maxValue)

	//Create a goroutines in loop with WaitGroup. And sending in channels message.
	for id := 0; id < amountOfThreads; id++ {
		wg.Add(1)
		go worker(ctx, cancel, c, &wg)
	}

	//Create goroutine for increment count.
	go c.Increment(&wgChan, cancel)
	wgChan.Add(1)
	wg.Wait()
	c.CloseChanel()
	wgChan.Wait()
	fmt.Println("Counter: ", c.Value())
}
