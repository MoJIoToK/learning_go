package main

import (
	"fmt"
	"regexp" // 1
)

func main() {
	re := regexp.MustCompile(`Hello, ([A-Z]{1}[a-z]+)`) // 2

	testStr := "test...^&* test test Hello, Foobar #$#$#$#$####" // 3

	fmt.Println(re.FindString(testStr)) // 4

	submatch := re.FindAllStringSubmatch(testStr, -1) // 5
	fmt.Println(submatch)
}

// Hello, Foobar
// Foobar
