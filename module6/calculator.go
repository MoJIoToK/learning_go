package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	for {
		fmt.Print("Введите первое число: ")
		fmt.Scan(&num1)

		fmt.Print("Введите желаемое действие: ")
		fmt.Scan(&operator)

		fmt.Print("Введите второе число: ")
		fmt.Scan(&num2)

		switch operator {
		case "+":
			fmt.Printf("Результат сложения: %v\n\n", num1+num2)
		case "-":
			fmt.Printf("Результат вычитания: %v\n\n", num1-num2)
		case "*":
			fmt.Printf("Результат умножения: %v\n\n", num1*num2)
		case "/":
			if num2 != 0 {
				fmt.Printf("Результат деления: %v\n\n", num1/num2)
			} else {
				panic("На ноль делить нельзя")
			}
		default:
			fmt.Println("Неизвестная операция! Попробуйте снова")
		}
	}
}
