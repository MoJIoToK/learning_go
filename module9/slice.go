package main

import "fmt"

func main() {
	//s1()

	//arr := [5]float64{1, 2, 3, 4, 5}
	//x := arr[0:5]
	//fmt.Println(arr, x)

	//sliceFromArray()

	//slice1 := []int{1, 2, 3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	var arr [5]int
	//sl := arr[:]
	sl := arr
	fmt.Println(sl)

}

//func s1() {
//	sl := make([]float64, 5)
//
//	fmt.Println(sl, cap(sl), len(sl))
//
//	sl = append(sl, 1)
//
//	fmt.Println()
//}

//func sliceFromArray() {
//	var arr = [5]int{1, 2, 3, 4, 5}
//	sl := arr[:]
//	fmt.Println(arr, sl)
//	sl[4] = 100
//	fmt.Println(arr, sl)
//}
