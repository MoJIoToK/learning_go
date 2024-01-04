package main

import "fmt"

func main() {
	var intArr = []int{1, 2, 4, 5, 6}
	var stringArr = []string{"one", "two", "three"}

	// элементом слайса может быть любой тип, в т.ч. слайс
	var matrix = [][]int{
		[]int{0, 1},
		[]int{3, 4},
	}
	fmt.Println(intArr)
	fmt.Println(stringArr)
	fmt.Println(matrix)
	firstElement := intArr[0]
	fmt.Println(firstElement)
	for i, value := range stringArr {
		fmt.Println(i, value)
	}
	matrix[1] = []int{5, 6}
	fmt.Println(matrix)

	ar := make([]int, 2, 4)
	fmt.Println(len(ar), cap(ar), ar)

	sar := make([]string, 2, 4)
	fmt.Println(len(sar), cap(sar), sar)

	par := make([]*struct{}, 2, 4)
	fmt.Println(len(par), cap(par), par)
}
