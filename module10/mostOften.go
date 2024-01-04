package main

import "fmt"

func main() {
	var slic = []int{1, -2, 4, -5, 6, -7, 8, 8, 10, 1, -4, 5, 10, -6, 10}
	fmt.Println(findMostOftenRepeated(slic))
	fmt.Println(findMostOftenRepeatedOptimized(slic))

}

// FindMostOftenRepeated returns most often repeat element in array
func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func findMostOftenRepeatedOptimized(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array[i+1:] {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}
