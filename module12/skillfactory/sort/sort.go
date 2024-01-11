package sort

import (
	"math/rand"
)

// region bubbleSort sorts array in ascending order. Smaller element in left, larger in right

func bubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
}

//endregion

// region selectionSort sorts array from left to right order
func selectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[minIndex] > ar[j] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

//endregion

// region insertionSort sorts array in ascending order. To the left of the current element is a sorted array
func insertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 1; i < len(ar); i++ {
		j := i
		for j > 0 {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
			j--
		}
	}
}

//endregion

// region mergeSort sorts array in ascending order. Divides the array into parts, compares them and
// combines already sorted arrays
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

//endregion

// region quickSort sorts array in ascending order. This sorts use pivot point
func quickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])

	return
}

//endregion
