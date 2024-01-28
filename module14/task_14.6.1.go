package main

import (
	"fmt"
)

// Intersection returns an array of type int with the same elements from two slices.
// Counts duplicate elements in the same slice. Parameters - two slices of type - int.
// Example, first slice - [1, 2, 3, 4, 5, 1], second slice - [3, 4, 5, 1], result slice - [3 4 5 1].
func intersection(inputArr1, inputArr2 []string) (result []string) {

	inputMap := make(map[string]int)
	for _, val := range inputArr1 {
		inputMap[val] = 1
	}

	for _, val := range inputArr2 {
		if _, ok := inputMap[val]; ok {
			inputMap[val]++
		}
	}

	for key, val := range inputMap {
		if val > 1 {
			result = append(result, key)
		}
	}
	return result
}

// FillArray fills the slice. Parameter - size of slice, int. Return - slice of int.
func fillArray(size int) (arr []string) {

	arr = make([]string, size)

	fmt.Println("Enter array:")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}

	return arr

}

func main() {
	var size1, size2 int

	fmt.Println("Enter first array size:")
	fmt.Scanln(&size1)

	fmt.Println("Enter second array size:")
	fmt.Scanln(&size2)

	arr1 := fillArray(size1)
	arr2 := fillArray(size2)

	fmt.Println("Intersection between two slices - ", intersection(arr1, arr2))

}
