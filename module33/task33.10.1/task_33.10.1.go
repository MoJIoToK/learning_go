package task33101

// CountLetters подсчитывает количество символов в строке.
func CountLetters(s string, l byte) int {
	c := 0
	for _, b := range []byte(s) {
		if b == l {
			c++
		}
	}
	return c
}
