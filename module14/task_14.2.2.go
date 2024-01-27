package main

import (
	"fmt"
)

// Hashstr returns a hash (of type int64) from strings (of type string), using the remainder of division by 1000.
func hashstr(s string) uint64 {
	hash := uint64(0)
	for _, char := range s {
		hash = hash*31 + uint64(char)
	}
	return hash % 1000
}

func main() {
	fmt.Println(hashstr("abc"))
	fmt.Println(hashstr("cba"))
	fmt.Println(hashstr("Hello"))
	fmt.Println(hashstr("World, hello"))

}
