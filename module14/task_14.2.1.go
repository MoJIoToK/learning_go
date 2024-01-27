package main

import "fmt"

// Hashint64 returns hash(type int64) of a input number(type int64) using the remainder of division by 1000.
func hashint64(val int64) (hash uint64) {
	hash = uint64(val % 1000)
	return hash
}

func main() {
	fmt.Println(hashint64(1))
	fmt.Println(hashint64(10))
	fmt.Println(hashint64(100))
	fmt.Println(hashint64(1000))

}
