package main

import (
	"fmt"
)

func main() {
	//i := 4
	//defer func(n int) {
	//	fmt.Println("Everything is fine:", n)
	//}(i)
	//i = 3
	//fmt.Println(i)

	Cascade()
	fmt.Println("Everything is fine")
}

// Вывод:
// 3
// Everything is fine: 4

func Cascade() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %v\n", r)
		}
	}()
	panicFunc()
	fmt.Println("Cascade end")
}

func panicFunc() {
	fmt.Println("just before panic")
	if true {
		panic("this is an artificial panic")
	}
	fmt.Println("panicFunc without panic")
}
