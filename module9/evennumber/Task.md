# Even number

Исправьте ошибки в программе:
```azure
package main

import (
	fmt "log"
)

// программа считает сумму чётных чисел от 0 до 50 включительно и выводит её
func main() int {
	var sum int64
	for i := 0; i < 50; i++ {
		sum += i
	}

	log.Println(sum)
}
```