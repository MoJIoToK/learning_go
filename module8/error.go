package main

import (
	"errors"
	"fmt"
	"os"
)

/*
func main() {
	i, err := someFunc(0)
	fmt.Println(i, err)

	i, err = someFunc(10)
	fmt.Println(i, err)
}

func someFunc(i int) (int, error) {
	if i == 0 {
		return 0, errors.New("got 0")
	}

	return i, nil
}
*/

func main() {
	i, err := someFunc(0)
	fmt.Println(i, err)

	i, err = someFunc(10)
	fmt.Println(i, err)

	testIs()

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

func someFunc(i int) (int, error) {
	j, err := funcReturningError(i)
	if err != nil {
		return 0, fmt.Errorf("wrap error: %w", err)
	}

	return j, nil
}

func funcReturningError(i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("got %d", i)
	}

	return i, nil
}

func testIs() {
	var errNotFound = errors.New("not found")

	err := fmt.Errorf("wrap: %w", errNotFound)

	if errors.Is(err, errNotFound) {
		fmt.Println("got error with Is")
	}
}
