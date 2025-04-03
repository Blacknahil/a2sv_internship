package utils

import (
	"clean_task_manager_api_tested/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDatabaseWrapper wraps *mongo.Database to implement domain.DatabaseInterface
type MongoDatabaseWrapper struct {
	database *mongo.Database
}

// NewMongoDatabaseWrapper creates a new wrapper for *mongo.Database
func NewMongoDatabaseWrapper(database *mongo.Database) domain.DatabaseInterface {
	return &MongoDatabaseWrapper{database: database}
}

// Collection implements the Collection method of domain.DatabaseInterface
func (w *MongoDatabaseWrapper) Collection(name string, opts ...*options.CollectionOptions) domain.CollectionInterface {
	collection := w.database.Collection(name, opts...)
	return NewMongoCollectionWrapper(collection)
}
