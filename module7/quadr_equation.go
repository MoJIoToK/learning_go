package main

import (
	"fmt"
	"math"
)

//написать функцию для нахождения корней квадратного уравнения через дискриминант.
//реализация принимает в качестве аргументов коэффициенты a, b, c, работает с вещественными
//числами (не работает с комплексными) и возвращает ошибки в случае возникновения
//ошибки объявим отдельно с помощью пакета `fmt`

var (
	//ErrZeroA выводится, когда уравнение не является квадратным.
	ErrZeroA = fmt.Errorf("coefficient a is zero")
	//ErrNoRealRoots выводится когда у уравнения нет вещественных корней.
	ErrNoRealRoots = fmt.Errorf("equation has no real roots")
)

// SolveQuadraticEquation finds real roots of equation defined with 3 real coefficients.
// It returns 2 roots if no error encountered or default float64 values and error otherwise.
func SolveQuadraticEquation(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 { // проверка на то, что уравнение квадратное
		err = ErrZeroA // возвращение ошибки
		// так как в сигнатуре даны имена переменным, то
		// возвращаем именно их; по умолчанию x1 = x2 = 0.0
		return
	}

	D := b*b - 4*a*c // вычисление дискриминанта
	if D < 0 {
		err = ErrNoRealRoots
		return
	}
	if D == 0 { // уравнение имеет два одинаковых корня
		x1 = -b / (2 * a)
		x2 = x1
		return // err == nil по умолчанию
	}

	dRoot := math.Sqrt(D)
	x1 = (-b + dRoot) / (2 * a)
	x2 = (-b - dRoot) / (2 * a)
	return
}

func main() {
	fmt.Println(SolveQuadraticEquation(1, 2, -3))

}
