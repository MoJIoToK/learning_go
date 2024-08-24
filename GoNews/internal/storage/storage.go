package storage

import (
	"GoNews/internal/model"
)

type DB interface {
	GetPosts() ([]model.Post, error)
	AddPost(post model.Post) (int, error)
	Close() error
}
