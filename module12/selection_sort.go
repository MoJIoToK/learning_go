package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 10)
	ar = []int{13, 6, 11, -5, 9, 0}

	//for i := range ar {
	//	ar[i] = rand.Intn(100) - 50 // ограничиваем случайное значение от [-100;100]
	//}
	fmt.Println(ar)

	//selectionSort(ar)
	//selectionSortByMax(ar)
	selectionSortBidirection(ar)

	fmt.Println(ar)
}

// SelectionSort sorts array from left to right order
func selectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[minIndex] > ar[j] {
				minIndex = j
				fmt.Println(i, j, ar)
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

// SelectionSort sorts array from right to left order
func selectionSortByMax(ar []int) {
	for i := len(ar) - 1; i >= 0; i-- {
		maxIndex := i
		for j := i - 1; j >= 0; j-- {
			if ar[maxIndex] < ar[j] {
				maxIndex = j
				fmt.Println(i, j, ar)
			}
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}

func selectionSortBidirection(ar []int) {
	maxPos := len(ar) - 1
	minPos := 0

	for minPos <= maxPos {
		var minIdx = minPos
		var maxIdx = maxPos
		for j := minPos; j <= maxPos; j++ {
			if ar[j] < ar[minIdx] {
				minIdx = j
			}
			if ar[j] > ar[maxIdx] {
				maxIdx = j
			}
		}
		ar[minPos], ar[minIdx] = ar[minIdx], ar[minPos]
		if maxIdx == minPos {
			maxIdx = minIdx
		}
		ar[maxPos], ar[maxIdx] = ar[maxIdx], ar[maxPos]

		minPos++
		maxPos--
	}
}
