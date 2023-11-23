package main

import "fmt"

func main() {

	//for i := 1; i < 100; i *= 2 {
	//	i += 4
	//
	//	if i%2 == 0 {
	//		println("чётное", i)
	//	}
	//}
	var i, j, count int
	//c := 0
	for i = 1; i < 10; i++ {
		for j = 9; j > 0; j-- {
			switch {
			default:
				println("Tuta")
				break
			case i == j:
				c := i + j
				c = c * 2
				println("Ne tuta")
				continue
			}
			count++
			println(i, j, count)
		}
	}
	fmt.Println(i * j)
}
