package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Напишите код, в котором имеются два канала сообщений из целых чисел так, чтобы приём сообщений всегда приводил к
// блокировке. Приёмом сообщений из обоих каналов будет заниматься главная горутина. Сделайте так, чтобы во время такого
// «бесконечного ожидания» сообщений выполнялась фоновая работа в виде вывода текущего времени в консоль.

// RandInt creates random int number from min to max.
func randInt1(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	//Sender is anonymous function. This function sends message to chanel.
	sender := func() {
		for {
			allocation := randInt(1, 100)
			//Send message with step from 1 second to 5 second. This will block the current goroutine. And `select`
			//always run to default way.
			<-time.Tick(time.Second * time.Duration(randInt(1, 5)))
			if allocation <= 50 {
				c1 <- randInt1(1, 100)
			} else {
				c2 <- randInt1(1, 100)
			}
		}
	}

	//Start goroutine for send message to channels.
	go sender()

	for {
		select {
		case <-c1:
			fmt.Println("Канал с1 принял сообщение: ")
		case <-c2:
			fmt.Println("Канал с2 принял сообщение: ")
		default:
			fmt.Println("Текущее время: ", time.Now().Format("15:04:05"))

		}
	}
}
