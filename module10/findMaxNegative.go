package main

import "fmt"

func main() {
	var slic = []int{1, -2, 4, -5, 6, -7, 8}

	fmt.Println(findMax(slic))
	fmt.Println(findMaxNegative(slic))
}

// FindMax returns max element in slice
func findMax(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	max = array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
	}

	return max, nil
}

// FindMaxNegative returns max negative element in slice
func findMaxNegative(array []int) (maxNegative int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	var negativeSlice = []int{}
	for _, val := range array {
		if val < 0 {
			negativeSlice = append(negativeSlice, val)
		}
	}
	fmt.Println(negativeSlice)
	if len(negativeSlice) == 0 {
		return 0, fmt.Errorf("could not found max in empty negative slice")
	}
	return findMax(negativeSlice)
}
