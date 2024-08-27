package mongo

import (
	"GoNews/internal/model"
	"GoNews/internal/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Константы для обращения к БД и коллекции в mongoDB
const (
	dbName         = "goNews"
	collectionName = "posts"
)

type Storage struct {
	db *mongo.Client
}

// New - конструктор подключения к БД
func New(conn string) (*Storage, error) {
	const operation = "storage.mongodb.new"

	mogngoOpts := options.Client().ApplyURI(conn)
	db, err := mongo.Connect(context.Background(), mogngoOpts)
	if err != nil {
		return nil, fmt.Errorf("%s %s", operation, err)
	}

	if err := db.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("%s %s", operation, err)
	}

	collection := db.Database(dbName).Collection(collectionName)
	//Создание индекса для поля ID структуры model.Post
	indexId := mongo.IndexModel{
		Keys:    bson.D{{Key: "title", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.Background(), indexId)
	if err != nil {
		return nil, fmt.Errorf("%s %s", operation, err)
	}

	return &Storage{db: db}, nil
}

// AddPost - записывает посты в БД. На вход метода подаётся слайс постов.
// Метод возвращает количество успешно записанных постов и ошибку, отличную от duplicate key error.
func (s *Storage) AddPost(posts []model.Post) (int, error) {
	const operation = "storage.mongodb.addPost"

	var inputDB []interface{}
	for _, post := range posts {
		bsn := bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "id", Value: post.ID},
			{Key: "title", Value: post.Title},
			{Key: "content", Value: post.Content},
			{Key: "pubTime", Value: post.PubTime},
			{Key: "link", Value: post.Link},
		}
		inputDB = append(inputDB, bsn)
	}

	collection := s.db.Database(dbName).Collection(collectionName)
	res, err := collection.InsertMany(context.Background(), inputDB)
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return len(res.InsertedIDs), fmt.Errorf("%s %s", operation, err)
	}
	return len(res.InsertedIDs), nil
}

// GetPosts - возвращает указанное число последних постов из БД.
// На вход принимается число публикаций, которое должно быть возвращено.
func (s *Storage) GetPosts(n int) ([]model.Post, error) {
	const operation = "storage.mongodb.getPosts"

	if n == 0 {
		return nil, fmt.Errorf("%s: %w", operation, storage.ErrZeroRequest)
	}

	collection := s.db.Database(dbName).Collection(collectionName)
	filter := bson.M{}

	opts := options.Find().SetLimit(int64(n))

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	var posts []model.Post
	err = cur.All(context.Background(), &posts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	if len(posts) == 0 {
		return nil, fmt.Errorf("%s: %w", operation, storage.ErrEmptyDB)
	}

	return posts, nil
}

// Close - обертка для закрытия пула подключений.
func (s *Storage) Close() error {
	return s.db.Disconnect(context.Background())
}
