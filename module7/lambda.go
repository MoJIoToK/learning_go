package main

import "fmt"

func main() {
	var c int
	func(a, b int) {
		c = a + b
	}(3, 5)
	fmt.Println("Hello, playground", c)
}
