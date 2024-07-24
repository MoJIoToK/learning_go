package tdd

import "math"

type Point struct {
	X, Y int
}

//Вычисление расстояние между двумя точками
func Distance(a, b Point) float64 {
	if a.X < 0 || b.X < 0 || a.Y < 0 || b.Y < 0 {
		return 0
	}
	res := math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
	return res
}
