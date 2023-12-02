package main

import (
	"fmt"
	elec "module8/task_8_8_1/electronic"
)

func main() {
	iPone := elec.NewApplePhone("first", "15")
	sasung := elec.NewAndroidPhone("second", "Samsung", "modelS")
	lg := elec.NewRadioPhone(12, "brand", "LG")

	printCharacteristics(iPone)
	printCharacteristics(sasung)
	printCharacteristics(lg)
}

func printCharacteristics(p elec.Phone) {
	if p.Type() == "station" {
		phone := p.(*elec.RadioPhone)
		fmt.Printf("Brand of this phone - %s, Model - %s, Type - %s, Count of Button - %d\n",
			phone.Brand(), phone.Model(), phone.Type(), phone.ButtonCount())
	} else {
		switch p.Brand() {
		case "Apple":
			phone := p.(*elec.ApplePhone)
			fmt.Printf("Brand of this phone - %s, Model - %s, Type - %s, OS - %s\n",
				phone.Brand(), phone.Model(), phone.Type(), phone.OS())
		default:
			phone := p.(*elec.AndroidPhone)
			fmt.Printf("Brand of this phone - %s, Model - %s, Type - %s, OS - %s\n",
				phone.Brand(), phone.Model(), phone.Type(), phone.OS())
		}
	}
}
