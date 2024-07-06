package main

//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//// сущность, которую храним в БД
//type book struct {
//	ID    int
//	Title string
//}
//
//func main() {
//	books := []book{
//		{ID: 1, Title: "1984"},
//		{ID: 2, Title: "Clean Architecture"},
//	}
//	// БД типа "ключ-значение"
//	kvstore := map[int]string{}
//	// наполнение БД
//	for _, b := range books {
//		bytes, _ := json.Marshal(b)
//		kvstore[b.ID] = string(bytes)
//	}
//	// печать содержимого БД
//	for k, v := range kvstore {
//		fmt.Printf("Ключ: %d, Значение: %s\n", k, v)
//	}
//}
