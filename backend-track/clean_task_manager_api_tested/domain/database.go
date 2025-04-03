package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database interface represents a generic database abstraction
// DatabaseInterface represents a generic database abstraction
type DatabaseInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) CollectionInterface
}

// CollectionInterface represents a generic MongoDB collection abstraction
type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}) SingleResultInterface
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (CursorInterface, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) // Added CountDocuments

}

// SingleResultInterface allows mocking FindOne operations
type SingleResultInterface interface {
	Decode(v interface{}) error
}

// CursorInterface allows mocking Find operations
type CursorInterface interface {
	Next(ctx context.Context) bool
	Decode(v interface{}) error
	Close(ctx context.Context) error
}
