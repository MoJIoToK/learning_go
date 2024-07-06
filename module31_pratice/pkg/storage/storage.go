package storage

import "module31_pratice/pkg/model"

// Interface - интерфейс для работы с БД
type Interface interface {
	GetPosts() ([]model.Post, error)
	AddPost(model.Post) (int, error)
	UpdatePost(int, model.Post) error
	DeletePost(id int) error
}
