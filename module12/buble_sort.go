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
	for i := range ar {
		ar[i] = rand.Intn(20) - 10 // ограничиваем случайное значение от [-100;100]
	}
	fmt.Println(ar)

	bubbleSort(ar)
	//bubbleSortWithInterrupt(ar)
	//bubbleSortReversed(ar)
	//bubbleSortRecursive(ar)
	fmt.Println(ar)

}

// BubbleSort sorts array in ascending order.
func bubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
}

// BubbleSortWithInterrupt sorts array in ascending order. If there is no swap during a complete search,
// the execution of the function was break.
func bubbleSortWithInterrupt(ar []int) {
	//swapped := true
	//for swapped {
	//	swapped = false
	//	for i := 0; i < len(ar)-1; i++ {
	//		if ar[i] > ar[i+1] {
	//			swapped = true
	//			ar[i], ar[i+1] = ar[i+1], ar[i]
	//		}
	//	}
	//}

	for i := 0; i < len(ar)-1; i++ {
		swapped := false
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				swapped = true
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSort sorts array in descending order.
func bubbleSortReversed(ar []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] < ar[i+1] {
				swapped = true
				ar[i], ar[i+1] = ar[i+1], ar[i]
			}
		}
	}
}

// BubbleSort sorts array in descending order with recursion. First, the largest element
// is placed at the end. Then the array is sent back to the function without the last element.
func bubbleSortRecursive(ar []int) {
	if len(ar) == 1 {
		return
	}
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i+1], ar[i] = ar[i], ar[i+1]
		}
	}
	fmt.Println(ar)
	bubbleSortRecursive(ar[:len(ar)-1])
}
