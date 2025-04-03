package infrastructure

import (
	"clean_task_manager_api_tested/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfAdmin(ctx context.Context, collection domain.CollectionInterface) (bool, error) {
	// check if the user collection is empty
	count, err := collection.CountDocuments(ctx, bson.M{})

	if err != nil {
		return false, err
	}
	// if so , return true
	if count == 0 {
		return true, nil
	}
	// else search for anyone who is already an admin
	filter := bson.M{"role": "admin"}

	adminCount, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return false, err
	}

	if adminCount > 0 {
		// return false
		return false, nil
	}

	return true, nil
	// return true to make the first user an admin
}
