package main

import (
	"fmt"
)

func main() {
	fmt.Println("Input string and press Enter")
	for {
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("error:", err)
			return
		}
		if input == "end" {
			return
		}
		fmt.Println(checkBrackets(input))
	}
}

func checkBrackets(input string) bool {
	stack := make([]rune, 0)
	for _, v := range input {
		switch v {
		case '{', '[', '(':
			stack = append(stack, v)
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (v == '}' && last != '{') || (v == ']' && last != '[') || (v == ')' && last != '(') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
