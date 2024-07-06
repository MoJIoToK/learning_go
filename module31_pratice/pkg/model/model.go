package model

// Post - модель публикации
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	CreatedAt   int64
	PublishedAt int64
}
