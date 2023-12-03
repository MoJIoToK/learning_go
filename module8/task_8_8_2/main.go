package main

import (
	"module8/task_8_8_2/mall"
)

func main() {
	//BMW
	bmw := mall.NewAuto(mall.CM, 500, 600, 700, "BMW", "X5",
		250, 250)
	merc := mall.NewAuto(mall.CM, 400, 500, 600, "Mercedes", "C",
		300, 500)
	dodge := mall.NewAuto(mall.Inch, 100, 200, 300, "Dodge", "Viper",
		250, 300)
	mall.Print(bmw, true)
	mall.Print(merc, true)
	mall.Print(dodge, false)

}
