package memdb

import "GoNews/internal/model"

type Storage struct {
	news []model.Post
}

// New - конструктор для эмулятора подключения к ДБ
func New() *Storage {
	return &Storage{}
}

// GetPosts - метод возвращает слайс с публикациями
func (s *Storage) GetPosts() ([]model.Post, error) {
	var posts []model.Post

	for _, post := range s.news {
		posts = append(posts, post)
	}

	return posts, nil
}

// AddPost - метод добавляет публикацию в память
func (s *Storage) AddPost(post model.Post) (int, error) {
	s.news = append(s.news, post)
	return len(s.news), nil
}

// Close - эмуляция закрытия БД
func (s *Storage) Close() error {
	return nil
}

// Len - метод возвращает количество публикаций
func (s *Storage) Len() int {
	return len(s.news)
}
