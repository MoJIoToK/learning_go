package storage

import (
	"GoNews/internal/model"
	"errors"
)

var (
	ErrEmptyDB     = errors.New("Database is empty")
	ErrZeroRequest = errors.New("Requested zero posts")
)

type DB interface {
	GetPosts(n int) ([]model.Post, error)
	AddPost(post []model.Post) (int, error)
	Close() error
}
