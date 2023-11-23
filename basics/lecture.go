package main

func main() {
	// var a int
	// var b int8
	// var c int16
	// var d int32
	// var e int64
	// var f float32
	// var g float64
	// var h complex64
	// var i complex128
	// var j rune
	// var k string
	// var l bool

	//fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
	//println(a, b)
	// Напишите программу, которая объявляет переменную типа int64. Имя переменной
	// укажите на свой вкус. Переменная должна инициализироваться целочисленным литералом
	// со значением 10 в шестнадцатеричном представлении. Затем программа объявляет вторую
	// переменную с типом int с инициализацией значением, хранящимся в первой переменной.
	// Имена переменных укажите на свой вкус.
	var a int64 = 0xA
	var b int = int(a)
	println(a, b)
}
