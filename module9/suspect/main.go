package main

import "fmt"

// Man is struct with human characteristics
type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

// Bd is a non struct type map[string]Man.
type bd map[string]Man

func main() {
	people := bd{
		"Tony":     {"Tony", "Soprano", 45, "male", 5},
		"Johny":    {"Johny", "Cash", 35, "male", 2},
		"Angelina": {"Angelina", "Jolie", 48, "female", 1},
		"Monica":   {"Monica", "Bellucci", 59, "female", 0},
		"Eren":     {"Eren", "Yeager", 16, "male", 3},
	}

	var suspects = []string{"Tony", "Monica", "Eren", "Bony"}
	var mostCriminalPerson Man
	var mostCriminalFound bool

	for _, susp := range suspects {
		person, ok := people[susp]
		if !ok {
			continue
		}
		if person.Crimes > mostCriminalPerson.Crimes {
			mostCriminalPerson = person
			mostCriminalFound = true
		}
	}

	if mostCriminalFound {
		fmt.Println("Наиболее криминальная личность -", mostCriminalPerson.Name, mostCriminalPerson.LastName)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}

}
