// djhsjd
package data

import (
	"context"
	"errors"
	"task_manager_api_auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskServices struct {
	TaskCollection *mongo.Collection
}

func (ts *TaskServices) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	task.DueDate = time.Now()

	result, err := ts.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	// change the insertedID to primitiveIdObj
	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}
	task.ID = objID
	// what if I just return the result
	return task, nil

}

func (ts *TaskServices) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := ts.TaskCollection.Find(ctx, bson.M{})
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

func (ts *TaskServices) GetTaskById(ctx context.Context, id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = ts.TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (ts *TaskServices) UpdateTask(ctx context.Context, id string, updatedTask models.Task) error {

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

	_, err = ts.TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskServices) DeleteTask(ctx context.Context, id string) error {

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.New("invalid task ID")
	}

	result, err := ts.TaskCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
