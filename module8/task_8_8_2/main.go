package main

import (
	"fmt"
	mall "module8/task_8_8_2"
)

func main() {
	//BMW
	bmw := mall.NewAuto(mall.CM, 500, 600, 700, "BMW", "X5", 250, 500)
	fmt.Println(bmw)

}
