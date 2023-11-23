package main

import "fmt"

func main() {
	//yearDetection()
	// dayDetection()
	//cicle()
	// while()
	//fmt.Println(factorial(-1), factorial(0), factorial(1), factorial(5), factorial(6))
	// personOperations()
	// figures()
	// listArrays()
	// maps()
	runes()

}

func dayDetection() {
	var day uint
	fmt.Println("Введите номер дня: ")
	fmt.Scan(&day)
	var dayText string
	switch day {
	case 1:
		dayText = "Понедельник"
	case 2:
		dayText = "Вторник"
	case 3:
		dayText = "Среда"
	case 4:
		dayText = "Четверг"
	case 5:
		dayText = "Пятница"
	case 6:
		dayText = "Суббота"
	case 7:
		dayText = "Воскресенье"
	default:
		dayText = "День введен неправильно"
	}

	fmt.Printf("День недели - %s", dayText)
}

func yearDetection() {
	var year uint
	fmt.Println("Введите год: ")
	fmt.Scan(&year)

	fmt.Println("Year is:", year)

	if year%400 == 0 || ((year%100 != 0) && (year%4 == 0)) {
		fmt.Printf("Год %d високосный", year)
	} else {
		fmt.Printf("Год %d невисокосный", year)
	}
}

func cicle() {
	count := 10
	for i := 0; i <= count; i++ {
		fmt.Print(i, " ")
	}
}

func while() {
	f := true
	n := 0
	for f {
		n++
		if n > 20 {
			f = false
		}
		fmt.Print(n, " ")
	}
}

func factorial(n int) int {
	if n < 0 {
		return -1
	}
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

type Person struct {
	Name    string
	Age     uint
	Country string
}

func (p *Person) Rename(name string) {
	p.Name = name
}

func (p Person) PrintName() {
	fmt.Println(p.Name)
}

func personOperations() {
	person := Person{
		"Vasya", 30, "Russia",
	}
	fmt.Println(person)
	fmt.Println(person.Name)
	person.Rename("Petya")
	person.PrintName()
}

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Perimetr() float64 {
	return t.A + t.B + t.C
}

func (r Rectangle) Perimetr() float64 {
	return (r.height + r.width) * 2
}

func (r Rectangle) Square() float64 {
	return r.height * r.width
}

func figures() {
	rect := Rectangle{4.5, 5.6}
	trian := Triangle{
		C: 5,
		B: 4,
		A: 3,
	}
	fmt.Println(rect.Perimetr(), rect.Square(), trian.Perimetr(), rect, trian)
}

func listArrays() {
	number := [5]int{1, 2, 3, 4, 5}
	fmt.Println(number)
	for i := 0; i < len(number); i++ {
		fmt.Print(number[i], " ")
	}

	number2 := []int{}
	for i := 0; i <= 10; i++ {
		number2 = append(number2, i)
		fmt.Println(number2, "Length:", len(number2), ": Capacity:", cap(number2))
	}

	number3 := make([]int, len(number2))
	copy(number3, number2)
	number3[2] = 100

	fmt.Println(number2)
	fmt.Println(number3)
}

func maps() {
	ages := map[string]int{
		"John": 25,
		"Mary": 18,
		"Bob":  80,
	}
	fmt.Println(ages["John"])
	delete(ages, "John")
	ages["John"] = 26
	age, ok := ages["John"]
	if !ok {
		fmt.Println("Name `John` not found")
	} else {
		fmt.Println(age)
	}
	fmt.Println(ages["John"])
}

func runes() {
	var r rune = 'Ы'
	fmt.Printf("%T\n", r)
	fmt.Println(string(r))

	message := "Привет, мир!"
	for _, c := range message {
		fmt.Printf("%c", c)
	}
}
