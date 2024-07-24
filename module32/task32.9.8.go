package main

func main() {
	n := factorial(3)
	if n != 6 {
		println("ошибка")
		return
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
