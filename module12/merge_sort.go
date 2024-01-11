package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}

	sorted := mergeSort(ar)

	fmt.Println(ar)
	fmt.Println(sorted)
}

func mergeSort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}

	middle := len(ar) / 2

	left, right := mergeSort(ar[:middle]), mergeSort(ar[middle:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
