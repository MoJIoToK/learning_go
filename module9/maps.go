package main

import "fmt"

func main() {
	//var x map[string]int
	//x := make(map[string]int)
	//x["key"] = 10
	//fmt.Println(x)

	//elements := make(map[string]string)
	//elements["H"] = "Hydrogen"
	//elements["He"] = "Helium"
	//elements["Li"] = "Lithium"
	//elements["Be"] = "Beryllium"
	//elements["B"] = "Boron"
	//elements["C"] = "Carbon"
	//elements["N"] = "Nitrogen"
	//elements["O"] = "Oxygen"
	//elements["F"] = "Fluorine"
	//elements["Ne"] = "Neon"

	//elements := make(map[string]string){
	//	"H": "Hydrogen",
	//	"He": "Helium",
	//	"Li": "Lithium",
	//	"Be": "Beryllium",
	//	"B": "Boron",
	//	"C": "Carbon",
	//	"N": "Nitrogen",
	//	"O": "Oxygen",
	//	"F": "Fluorine",
	//	"Ne": "Neon",
	//}

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
	}

	//println(elements["Li"])
	//name, ok := elements["Un"]
	//println(name, ok)

	for key, val := range elements {
		fmt.Println(key, val)
	}
}
