package main

import "fmt"

// программа считает сумму чётных чисел от 0 до 50 включительно и выводит её
func main() {
	var sum int
	for i := 0; i < 50; i += 2 {
		sum += i
	}

	fmt.Println(sum)
}
