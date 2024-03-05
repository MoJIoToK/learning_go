package main

import (
	"fmt"
	"sync"
)

//Напишите программу, которая делает следующее: одна горутина по порядку отсылает числа от 1 до 100 в канал,
//вторая горутина их принимает в правильном порядке и печатает на экран (в консоль).

func main() {

	c1 := make(chan int)
	var wg sync.WaitGroup

	//Sender is anonymous function. This function sends message to chanel.
	sender := func() {
		for i := 0; i <= 100; i++ {
			c1 <- i
		}
		//Closing the channel for recording.
		close(c1)
		//Termination one goroutine in the waitgroup.
		wg.Done()
	}

	//Receiver is anonymous function. This function receives message from chanel.
	receiver := func() {
		defer wg.Done()
		for num := range c1 {
			fmt.Println(num)
		}
	}
	//Start goroutine for send message to channels.
	go sender()

	//Start goroutine for receive message from channels.
	go receiver()

	//Add counter of waitgroup to count of goroutine.
	wg.Add(2)
	wg.Wait()
}
