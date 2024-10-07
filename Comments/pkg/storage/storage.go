package storage

import (
	"Comments/pkg/model"
	"errors"
)

var (
	ErrEmptyDB     = errors.New("Database is empty")
	ErrZeroRequest = errors.New("Requested zero posts")
)

type DB interface {
	AddComment(comment model.Comment) error
	Comments(news string) ([]model.Comment, error)
}
