package utils

import (
	"context"

	"clean_task_manager_api_tested/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCollectionWrapper wraps *mongo.Collection to implement domain.CollectionInterface
type MongoCollectionWrapper struct {
	collection *mongo.Collection
}

// NewMongoCollectionWrapper creates a new wrapper for *mongo.Collection
func NewMongoCollectionWrapper(collection *mongo.Collection) domain.CollectionInterface {
	return &MongoCollectionWrapper{collection: collection}
}

// InsertOne implements the InsertOne method of domain.CollectionInterface
func (w *MongoCollectionWrapper) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return w.collection.InsertOne(ctx, document)
}

// FindOne implements the FindOne method of domain.CollectionInterface
func (w *MongoCollectionWrapper) FindOne(ctx context.Context, filter interface{}) domain.SingleResultInterface {
	return &MongoSingleResultWrapper{result: w.collection.FindOne(ctx, filter)}
}

// UpdateOne implements the UpdateOne method of domain.CollectionInterface
func (w *MongoCollectionWrapper) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return w.collection.UpdateOne(ctx, filter, update, opts...)
}

// DeleteOne implements the DeleteOne method of domain.CollectionInterface
func (w *MongoCollectionWrapper) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return w.collection.DeleteOne(ctx, filter, opts...)
}

// Find implements the Find method of domain.CollectionInterface
func (w *MongoCollectionWrapper) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (domain.CursorInterface, error) {
	cursor, err := w.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return &MongoCursorWrapper{cursor: cursor}, nil
}

// CountDocuments implements the CountDocuments method of domain.CollectionInterface
func (w *MongoCollectionWrapper) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return w.collection.CountDocuments(ctx, filter, opts...)
}
