package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoCursorWrapper wraps *mongo.Cursor to implement domain.CursorInterface
type MongoCursorWrapper struct {
	cursor *mongo.Cursor
}

// Next implements the Next method of domain.CursorInterface
func (w *MongoCursorWrapper) Next(ctx context.Context) bool {
	return w.cursor.Next(ctx)
}

// Decode implements the Decode method of domain.CursorInterface
func (w *MongoCursorWrapper) Decode(v interface{}) error {
	return w.cursor.Decode(v)
}

// Close implements the Close method of domain.CursorInterface
func (w *MongoCursorWrapper) Close(ctx context.Context) error {
	return w.cursor.Close(ctx)
}
