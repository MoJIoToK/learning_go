package main

import (
	"encoding/json"
	"fmt"
)

// Модель данных для товара.
type product struct {
	ID    int
	Name  string
	Price int
}

// Товары из БД.
var products = []product{
	{ID: 1, Name: "Шоколадка 'Алёнка'", Price: 100},
	{ID: 2, Name: "Шоколадка 'Сникерс'", Price: 150},
}

// Каталог товаров в памяти.
var cache = make(map[int]string)

func main() {
	// Для каждого товара получаем его представление
	// в формате JSON и записываем его в хэш-таблицу.
	for _, p := range products {
		b, _ := json.Marshal(p)
		cache[p.ID] = string(b)
	}
	fmt.Print(cache[1], "\n", cache[2])
}
