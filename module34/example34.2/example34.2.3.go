package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите имя: ")

	buf := make([]byte, 255)
	_, _ = os.Stdin.Read(buf)

	fmt.Printf("Привет, %s", buf)
}
