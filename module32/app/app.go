package main

import (
	"fmt"
	"github.com/DmitriyVTitov/size"
)

func main() {
	object := struct {
		i int
		s string
		m map[int]string
	}{
		i: 10,
		s: "string",
		m: map[int]string{
			10: "A",
			20: "B",
		},
	}
	fmt.Println(size.Of(object))

}
