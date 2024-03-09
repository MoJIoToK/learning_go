package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	countOfGoroutines := 0
	maxValue := 0

	//Create a channel.
	ch := make(chan int)

	//Create a syncgroup for increment count and channel filling.
	chanWg := sync.WaitGroup{}
	wg := sync.WaitGroup{}

	fmt.Println("Укажите количество горутин: ")
	_, err := fmt.Scanln(&countOfGoroutines)
	if err != nil {
		fmt.Println("Неверное значние")
	}
	if countOfGoroutines < 1 {
		fmt.Println("Количество горутин не может быть меньше 1")
	}

	fmt.Println("Укажите максимальное значение счётчика: ")
	_, err = fmt.Scanln(&maxValue)
	if err != nil {
		fmt.Println("Неверное значние")
	}
	if maxValue < 1 {
		fmt.Println("Счётчик не может быть меньше 1")
	}

	//Create goroutine in waitgroup for increment count.
	chanWg.Add(1)
	go func() {
		for range ch {
			if count < maxValue {
				count++
			}
		}
		chanWg.Done()
	}()

	//Create goroutine in waitgroup. Their number is equal to the number of goroutines.
	wg.Add(countOfGoroutines)
	//Creating Goroutines in a Loop.
	for i := 0; i < countOfGoroutines; i++ {
		go func(i int) {
			//Recording numbers in the channel
			ch <- i
			//Signal the end of one goroutine in a group.
			wg.Done()
		}(i)
	}
	//Wait all goroutine in group.
	wg.Wait()
	//Сlosing the recording channel
	close(ch)
	//Wait all goroutine in group.
	chanWg.Wait()

	fmt.Println("Конечное значение счётчика равно:", count)

}
