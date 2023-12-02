package main

import "fmt"

func main() {
	printInterfaceVal(10)
	printInterfaceVal("string")
	printInterfaceVal(true)
}

func printInterfaceVal(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("целое число: %d\n", v)
	case string:
		fmt.Printf("строка: %s\n", v)
	case bool:
		fmt.Printf("булев тип: %t\n", v)
	}
}
