package memdb

import "module31_pratice/pkg/model"

// Store - структура хранилища данных
type Store struct {
}

// New - Конструктор БД основанной на памяти компьютера
func New() *Store {
	return new(Store)
}

// GetPosts - метод возвращающий список публикаций и ошибку
func (s *Store) GetPosts() ([]model.Post, error) {
	return posts, nil
}

// AddPost - метод добавляющий публикацию.
// Метод принимает на вход структуру model.Post.
// Метод возвращает ID публикации и ошибку
func (s *Store) AddPost(post model.Post) (int, error) {
	return 0, nil
}

// UpdatePost - метод обновляет публикацию по ID.
// Метод принимает на вход ID и структуру model.Post.
// Метод возвращает ошибку
func (s *Store) UpdatePost(id int, post model.Post) error {
	return nil
}

// DeletePost - метод удаляет публикацию по ID.
// Метод принимает на вход ID.
// Метод возвращает ошибку
func (s *Store) DeletePost(id int) error {

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
		}
	}

	return nil
}

// Слайс, хранящий публикации
var posts = []model.Post{
	{
		ID:      1,
		Title:   "Effective Go",
		Content: "Go is a new language.",
	},
	{
		ID:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which",
	},
}
