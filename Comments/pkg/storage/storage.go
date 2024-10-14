package storage

import (
	"Comments/pkg/model"
	"context"
	"errors"
)

var (
	ErrNoComments         = errors.New("No comments on provided post id")
	ErrParentNotFound     = errors.New("Parent comment not found")
	ErrIncorrectParentID  = errors.New("Incorrect parent id")
	ErrIncorrectPostID    = errors.New("Incorrect post id")
	ErrIncorrectCommentID = errors.New("Incorrect comment id")
	ErrEmptyContent       = errors.New("Empty comment content field")
)

type DB interface {
	AddComment(ctx context.Context, comment model.Comment) (string, error)
	Comments(ctx context.Context, news string) ([]model.Comment, error)
}
