package main

import (
	"encoding/json"
	"fmt"
)

type Taged struct {
	FieldA int    `json:"field_a"`
	FieldB string `json:"field_b"`
	fieldC string `json:"field_c"`
}

type C struct {
	d int
}

type B struct {
	c C
}

type A struct {
	b B
}

func main() {
	c := C{d: 5}
	b := B{c: c}
	a := A{b: b}

	fmt.Printf("%+v\n", a)

	someJsonContent := `
{
"field_a": 15,
"field_b": "Это строка будет записана в поле!",
"field_c": "Это строка не будет записана!"
}
`
	taggedStruct := Taged{}
	if err := json.Unmarshal([]byte(someJsonContent), &taggedStruct); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", taggedStruct)

}
