package example

// Fact возвращает факториал числа.
func Fact(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i
	}
	return f
}
