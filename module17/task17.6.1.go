package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// RandInt creates random int number from min to max.
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	var countC1 float64 = 0
	var countC2 float64 = 0
	var wg sync.WaitGroup

	//Sender is anonymous function. This function sends message to chanel. First chanel gets more message.
	sender := func() {
		for {
			allocation := randInt(1, 100)
			if allocation <= 50 {
				c1 <- randInt(1, 100)
			} else {
				allocation = randInt(1, 100)
				if allocation <= 50 {
					c2 <- randInt(1, 100)
				} else {
					continue
				}
			}
		}
	}

	//Receiver is anonymous function. This function receives message from chanel.
	receiver := func() {
		for {
			select {
			case num := <-c1:
				{
					countC1++
					var ratio = (countC1 / (countC1 + countC2)) * 100
					str := fmt.Sprintf("Канал с1 принимает %.2f%% сообщений. Сообщение из канала: %d", ratio, num)
					fmt.Println(str)
				}
			case num := <-c2:
				{
					countC2++
					var ratio = (countC2 / (countC1 + countC2)) * 100
					str := fmt.Sprintf("Канал с2 принимает %.2f%% сообщений. Cообщение из канала: %d", ratio, num)
					fmt.Println(str)
				}
			default:
				fmt.Println("Сообщений нет!")
			}
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
