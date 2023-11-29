package main

import "fmt"

func main() {
	args := []int{1, 2, 3}
	fmt.Println(SumPolynomial(2, args...)) // 34
	fmt.Println(SumPolynomial(3, 1, 2))    // 21
	fmt.Println(SumPolynomial(4))          // 0
	fmt.Println(SumPolynomial(5, 6))       // 30
}
func SumPolynomial(x float64, a ...int) int {
	pow := 1.0
	val := 0.0
	for i := 0; i < len(a); i++ {
		pow *= x
		val += float64(a[i]) * pow
		fmt.Println(pow, val)
	}
	return int(val)
}
