package main

import (
	"fmt"
	"testing/mock/pkg/model"
	"testing/mock/pkg/storage"
)

var db storage.Interface

func main() {
	db = postgres.New()
	t := task(db, 0)
	fmt.Println(t)
}

func task(db storage.Interface, id int) model.Task {
	tt, err := db.Tasks()
	if err != nil {
		return model.Task{}
	}
	for _, t := range tt {
		if t.ID == id {
			return t
		}
	}
	return model.Task{}
}
