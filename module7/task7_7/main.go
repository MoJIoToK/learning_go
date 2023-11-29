package main

import (
	"fmt"
	"module7/task7_7/calc"
)

func main() {
	var num1, num2, result float64
	var operator string
	var err error

	calculator := calc.NewCalculator()

	for {
		fmt.Print("Введите первое число: ")
		_, err = fmt.Scan(&num1)
		if err != nil {
			fmt.Printf("Ошибка при чтении первого числа %v", err)
		}

		fmt.Print("Введите желаемое действие: ")
		_, err = fmt.Scan(&operator)
		if err != nil {
			fmt.Printf("Ошибка при чтении оператора %v", err)
		}

		fmt.Print("Введите второе число: ")
		_, err = fmt.Scan(&num2)
		if err != nil {
			fmt.Printf("Ошибка при чтении второго числа %v", err)
		}

		result = calculator.Calculate(num1, num2, operator)

		fmt.Printf("Результат вычислений: %v\n"+
			"=============================\n", result)

	}
}
