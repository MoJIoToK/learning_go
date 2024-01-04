package main

import "fmt"

func Top(t func() int) int {
	return t()
}

func main() {
	a := 2
	b := &a
	mid := func() int {
		return a * *b * Top(func() int {
			a = 2
			return 3
		})
	}
	c := Top(mid)
	b = &c

	fmt.Println(*b)
}
