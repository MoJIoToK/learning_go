package storage

import (
	"GoNews/internal/model"
)

type Storage interface {
	GetPosts() ([]model.Post, error)
	AddPost(post model.Post) (int, error)
	Close() error
}
