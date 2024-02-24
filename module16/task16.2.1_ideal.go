package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	var number int

	fmt.Println("Enter the number of the goroutine")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
		return
	}

	//цикл для активации горутин
	for i := 0; i < number; i++ {
		go func() {
			//бесконечный цикл для печати идентификатора горутин
			for {
				fmt.Println("This is goroutine with id - ", i)
				time.Sleep(time.Second)
			}

		}()
	}
	//Задержка выполенения основной горутины. Для того чтобы вспомогательные смогли написать хоть что-то
	time.Sleep(time.Second * 30)

}
