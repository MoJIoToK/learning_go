package calc

import "fmt"

const (
	operatorAddition       = "+"
	operatorSubtraction    = "-"
	operatorMultiplication = "*"
	operatorDivision       = "/"
)

type calculator struct {
}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) Calculate(num1, num2 float64, operator string) (result float64) {
	switch operator {
	case operatorAddition:
		result = c.addition(num1, num2)
	case operatorSubtraction:
		result = subtraction(num1, num2)
	case operatorMultiplication:
		result = multiplication(num1, num2)
	case operatorDivision:
		result = division(num1, num2)
	default:
		fmt.Println("Неизвестная операция! Попробуйте снова")
	}

	return result
}

func (c *calculator) addition(summand1, summand2 float64) float64 {
	return summand1 + summand2
}

func subtraction(minued, subtrahend float64) float64 {
	return minued - subtrahend
}

func multiplication(multiplier, multiplicand float64) float64 {
	return multiplier * multiplicand
}

func division(numerator, denumerator float64) float64 {
	if denumerator != 0 {
		return numerator / denumerator
	} else {
		panic("На ноль делить нельзя")
	}
}
