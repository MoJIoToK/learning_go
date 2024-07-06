package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"module31_pratice/pkg/model"
)

// Storage - структура хранилища данных
type Storage struct {
	db *pgxpool.Pool
}

// New - конструктор для БД на основе postgresql
func New(constr string) (*Storage, error) {
	db, err := pgxpool.New(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

// GetPosts - метод возвращающий список публикаций и ошибку
func (store *Storage) GetPosts() ([]model.Post, error) {

	query := "SELECT * FROM posts"
	var posts []model.Post

	rows, err := store.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	//Запись данных из БД в слайс
	for rows.Next() {
		var post model.Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.PublishedAt,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, rows.Err()
}

// AddPost - метод добавляющий публикацию.
// Метод принимает на вход структуру model.Post.
// Из структуры model.Post берутся поля title, content и author_id, остальные заполняются автоматически.
// Метод возвращает ID публикации и ошибку
func (store *Storage) AddPost(post model.Post) (int, error) {
	var id int
	query := "INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id"
	err := store.db.QueryRow(context.Background(), query, post.Title, post.Content, post.AuthorID).Scan(&id)
	return id, err
}

// UpdatePost - метод обновляет публикацию по ID.
// Метод принимает на вход ID и структуру model.Post.
// Обновлению подлежат поля title, content и author_id структуры model.Post.
// Метод возвращает ошибку.
func (store *Storage) UpdatePost(id int, post model.Post) error {
	query := "UPDATE posts SET title = $1, content = $2, author_id = $3 WHERE id = $4"
	_, err := store.db.Exec(context.Background(), query, post.Title, post.Content, post.AuthorID, id)
	return err
}

// DeletePost - метод удаляет публикацию по ID.
// Метод принимает на вход ID.
// Метод возвращает ошибку
func (store *Storage) DeletePost(id int) error {
	query := "DELETE FROM posts WHERE id = $1"
	_, err := store.db.Exec(context.Background(), query, id)
	return err
}
