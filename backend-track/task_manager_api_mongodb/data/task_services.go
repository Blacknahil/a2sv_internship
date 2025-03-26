package data

import (
	"context"
	"errors"
	"task_manager_api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	Collection mongo.Collection
}

func (ts *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {

	var tasks []models.Task
	cursor, err := ts.Collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var task models.Task

		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (ts *TaskService) GetTaskById(ctx context.Context, id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return nil, errors.New("invalid task ID")
	}

	var task models.Task
	err = ts.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (ts *TaskService) UpdateTask(ctx context.Context, id string, updatedTask models.Task) error {

	// validate Id
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	updatedFields_bson := bson.M{}

	if updatedTask.Title != "" {
		updatedFields_bson["title"] = updatedTask.Title
	}

	if updatedTask.Description != "" {
		updatedFields_bson["description"] = updatedTask.Description
	}

	if updatedTask.Status != "" {
		updatedFields_bson["status"] = updatedTask.Status
	}

	if !updatedTask.DueDate.IsZero() {
		updatedFields_bson["due_date"] = updatedTask.DueDate
	}

	if len(updatedFields_bson) == 0 {

		return errors.New("no valid fields to update")
	}

	update := bson.M{"$set": updatedFields_bson}

	_, err = ts.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)

	if err != nil {
		return err
	}

	return nil

}

func (ts *TaskService) DeleteTask(ctx context.Context, id string) error {

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.New("invalid task ID")
	}

	result, err := ts.Collection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil

}

func (ts *TaskService) CreateTask(ctx context.Context, newTask *models.Task) (*models.Task, error) {

	// newTask.ID = primitive.NewObjectID().Hex()
	newTask.DueDate = time.Now()

	result, err := ts.Collection.InsertOne(ctx, newTask)
	if err != nil {
		return nil, err
	}
	// Type assert the InsertedID to primitive.ObjectID
	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	// Convert the ObjectID to a string and assign it to newTask.ID
	newTask.ID = objID
	return newTask, nil

}

// djbskjdbkjds
// dfhjshjdjhsd
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// fdjh jhfsdjhsjhdjh/jkdfbjkgkdfbjdfdjhsdhjhjsdhjjkd
// dhjfsdjhhj
// dhjsdjhjhsdjhhfhjdj
// fbdjbafdjaj / hjsdkfnjsdkbdhjjhsjsjdjdsjsdjjjshjhjsdhhshsjhjhshjshjhhjhdjjh
