package main

import (
	"fmt"
	"testing/statistics/pkg/statistics"
)

func main() {
	//Верный результат
	nums := []float64{1, 2, 3}
	want := 2.0
	got := statistics.Avg(nums)
	if got != want {
		fmt.Println("Получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)

	// ошибка
	nums = []float64{1, 2, 3}
	want = 3.0
	got = statistics.Avg(nums)
	if got != want {
		fmt.Printf("получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)

	// пустой массив
	nums = []float64{}
	want = 0
	got = statistics.Avg(nums)
	if got != want {
		fmt.Printf("получено %f, ожидалось %f\n", got, want)
	}
	fmt.Printf("результат: %f\n", got)
}
