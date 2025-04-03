package utils

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoSingleResultWrapper wraps *mongo.SingleResult to implement domain.SingleResultInterface
type MongoSingleResultWrapper struct {
	result *mongo.SingleResult
}

// Decode implements the Decode method of domain.SingleResultInterface
func (w *MongoSingleResultWrapper) Decode(v interface{}) error {
	return w.result.Decode(v)
}
