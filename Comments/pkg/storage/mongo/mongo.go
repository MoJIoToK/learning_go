package mongo

import (
	"Comments/pkg/model"
	"Comments/pkg/storage"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *Storage) AddComment(ctx context.Context, comment model.Comment) (string, error) {
	const operation = "storage.mongodb.addComment"

	if t, err := primitive.ObjectIDFromHex(comment.NewsID); err != nil {
		fmt.Println(t, err)
		return "", fmt.Errorf("%s: %w", operation, storage.ErrIncorrectPostID)
	}

	/*
		if comment.NewsID == "" {
			return fmt.Errorf("%s: %v", operation, storage.ErrEmptyDB)
		}

		if comment.Content == "" {
			return fmt.Errorf("%s: %v", operation, storage.ErrEmptyDB)
		}
	*/

	id := primitive.NewObjectID()
	bsn := bson.D{
		{Key: "_id", Value: id},
		{Key: "ParentID", Value: comment.ParentID},
		{Key: "NewsID", Value: comment.NewsID},
		{Key: "PubTime", Value: primitive.NewDateTimeFromTime(time.Now())},
		{Key: "Content", Value: comment.Content},
		{Key: "Childs", Value: bson.A{}},
	}

	collection := s.db.Database(dbName).Collection(collectionName)

	if comment.ParentID != "" {
		id, err := primitive.ObjectIDFromHex(comment.ParentID)
		if err != nil {
			return "", fmt.Errorf("%s: %w", operation, storage.ErrIncorrectParentID)
		}
		filter := bson.D{
			{Key: "_id", Value: id},
		}
		res := collection.FindOne(ctx, filter)
		if res.Err() != nil {
			if res.Err() == mongo.ErrNoDocuments {
				return "", fmt.Errorf("%s: %w", operation, storage.ErrParentNotFound)
			}
			return "", fmt.Errorf("%s: %w", operation, res.Err())
		}

	}

	_, err := collection.InsertOne(ctx, bsn)
	if err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	return id.Hex(), nil
}

func (s *Storage) Comments(ctx context.Context, news string) ([]model.Comment, error) {
	const operation = "storage.mongodb.AddComment"

	if news == "" {
		return nil, storage.ErrIncorrectPostID
	}

	if _, err := primitive.ObjectIDFromHex(news); err != nil {
		return nil, fmt.Errorf("%s: %w", operation, storage.ErrIncorrectPostID)
	}

	var comments []model.Comment
	collection := s.db.Database(dbName).Collection(collectionName)
	opts := options.Find().SetSort(bson.D{{Key: "PubTime", Value: -1}})
	filter := bson.D{{Key: "NewsID", Value: news}}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	if len(comments) == 0 {
		return nil, fmt.Errorf("%s: %w", operation, storage.ErrNoComments)
	}
	return comments, nil
}
