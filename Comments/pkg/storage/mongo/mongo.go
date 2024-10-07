package mongo

import (
	"Comments/pkg/model"
	"Comments/pkg/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	dbName         = "goComments"
	collectionName = "comments"
)

type Storage struct {
	db *mongo.Client
}

func New(conn string) (*Storage, error) {
	const operation = "storage.mongodb.new"

	mongoOpts := options.Client().ApplyURI(conn)
	db, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("%s %s", operation, err)
	}

	if err := db.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("%s %s", operation, err)
	}

	collection := db.Database(dbName).Collection(collectionName)

	indexId := mongo.IndexModel{
		Keys: bson.D{
			{
				Key: "NewsID", Value: -1,
			},
		},
	}

	_, err = collection.Indexes().CreateOne(context.Background(), indexId)

	if err != nil {
		return nil, fmt.Errorf("%s %s", operation, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddComment(comment model.Comment) error {
	const operation = "storage.mongodb.addComment"

	if comment.NewsID == "" {
		return fmt.Errorf("%s: %v", operation, storage.ErrEmptyDB)
	}

	if comment.Content == "" {
		return fmt.Errorf("%s: %v", operation, storage.ErrEmptyDB)
	}

	bsn := bson.D{
		{Key: "_id", Value: primitive.NewObjectID()},
		{Key: "ParentID", Value: comment.ParentID},
		{Key: "NewsID", Value: comment.NewsID},
		{Key: "PubTime", Value: primitive.NewDateTimeFromTime(time.Now())},
		{Key: "allowed", Value: true},
		{Key: "Content", Value: comment.Content},
		{Key: "Childs", Value: bson.A{}},
	}

	collection := s.db.Database(dbName).Collection(collectionName)

	if comment.ParentID == "" {
		_, err := collection.InsertOne(context.Background(), bsn)
		if err != nil {
			return fmt.Errorf("%s: %v", operation, storage.ErrEmptyDB)
		}
		fmt.Println(comment.ID)
		return nil
	}

	parent, err := primitive.ObjectIDFromHex(comment.ParentID)

	if err != nil {
		return fmt.Errorf("%s: %w", operation, storage.ErrEmptyDB)
	}
	opts := options.Update().SetUpsert(false)
	filter := bson.D{
		{Key: "NewsID", Value: comment.NewsID},
		{Key: "_id", Value: parent},
	}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "Childs", Value: bsn},
		}},
	}
	result, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("%s: %w", operation, storage.ErrEmptyDB)
	}
	return nil
}

func (s *Storage) Comments(news string) ([]model.Comment, error) {
	const operation = "storage.mongodb.AddComment"
	if news == "" {
		return nil, storage.ErrEmptyDB
	}
	var comments []model.Comment
	collection := s.db.Database(dbName).Collection(collectionName)
	opts := options.Find().SetSort(bson.D{{Key: "PubTime", Value: -1}})
	filter := bson.D{{Key: "NewsID", Value: news}}
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	err = cursor.All(context.Background(), &comments)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	if len(comments) == 0 {
		return nil, storage.ErrEmptyDB
	}
	return comments, nil
}
