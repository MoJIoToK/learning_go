package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"module31_pratice/pkg/model"
)

// Константы для обращения к БД и коллекции в mongoDB
const (
	dbName         = "goNews"
	collectionName = "posts"
)

// Storage - структура хранилища данных
type Storage struct {
	db *mongo.Client
}

// New - конструктор для БД на основе mongoDB
func New(conn string) (*Storage, error) {
	// создание клиента для работы с бд
	mogngoOpts := options.Client().ApplyURI(conn)
	//Подключение к БД
	db, err := mongo.Connect(context.Background(), mogngoOpts)
	if err != nil {
		return nil, err
	}

	//Проверка подключения к бд
	if err := db.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	//Создание индекса для поля ID структуры model.Post
	indexId := mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	collection := db.Database(dbName).Collection(collectionName)
	_, err = collection.Indexes().CreateOne(context.Background(), indexId)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

// GetPosts - метод возвращающий список публикаций и ошибку
func (s *Storage) GetPosts() ([]model.Post, error) {
	collection := s.db.Database(dbName).Collection(collectionName)
	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var posts []model.Post

	for cur.Next(context.Background()) {
		var post model.Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// AddPost - метод добавляющий публикацию.
// Метод принимает на вход структуру model.Post.
// Cтруктура model.Post записывается в БД как документ в
// формате JSON с присвоением _ID учтенной непосредственно в БД.
// Метод возвращает ID публикации и ошибку
func (s *Storage) AddPost(post model.Post) (int, error) {
	collection := s.db.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return 0, err
	}
	return post.ID, nil
}

// UpdatePost - метод обновляет публикацию по ID.
// Метод принимает на вход ID и структуру model.Post.
// Обновлению подлежат поля title, content и authorid структуры model.Post.
// Метод возвращает ошибку.
func (s *Storage) UpdatePost(id int, post model.Post) error {
	collection := s.db.Database(dbName).Collection(collectionName)

	filter := bson.D{{Key: "id", Value: id}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{"title", post.Title},
			{"content", post.Content},
			{"authorid", post.AuthorID},
		}},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// DeletePost - метод удаляет публикацию по ID.
// Метод принимает на вход ID.
// Метод возвращает ошибку
func (s *Storage) DeletePost(id int) error {
	collection := s.db.Database(dbName).Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return err
}
