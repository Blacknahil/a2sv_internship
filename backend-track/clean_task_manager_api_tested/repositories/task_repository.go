// dkkjdjkf
package repository

import (
	"clean_task_manager_api_tested/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepositoryImpl struct {
	database   domain.DatabaseInterface
	collection string
}

func (tr *taskRepositoryImpl) Create(c context.Context, task *domain.Task) (*domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	result, err := collection.InsertOne(c, task)
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

func (tr *taskRepositoryImpl) GetAllTasks(c context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	collection := tr.database.Collection(tr.collection)
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {

		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *taskRepositoryImpl) GetTaskByID(c context.Context, taskID string) (*domain.Task, error) {

	var task domain.Task
	collection := tr.database.Collection(tr.collection)

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (tr *taskRepositoryImpl) DeleteTask(c context.Context, taskID string) error {

	objId, err := primitive.ObjectIDFromHex(taskID)

	if err != nil {
		return errors.New("invalid task ID")
	}
	collection := tr.database.Collection(tr.collection)

	result, err := collection.DeleteOne(c, bson.M{"_id": objId})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (tr *taskRepositoryImpl) UpdateTask(c context.Context, taskID string, task *domain.Task) error {
	// validate Id
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return errors.New("invalid task ID")
	}

	updatedFields_bson := bson.M{}
	collection := tr.database.Collection(tr.collection)

	if task.Title != "" {
		updatedFields_bson["title"] = task.Title
	}

	if task.Description != "" {
		updatedFields_bson["description"] = task.Description
	}

	if task.Status != "" {
		updatedFields_bson["status"] = task.Status
	}

	if !task.DueDate.IsZero() {
		updatedFields_bson["due_date"] = task.DueDate
	}

	if len(updatedFields_bson) == 0 {

		return errors.New("no valid fields to update")
	}

	update := bson.M{"$set": updatedFields_bson}

	_, err = collection.UpdateOne(c, bson.M{"_id": objID}, update)

	if err != nil {
		return err
	}

	return nil
}

func NewTaskRepositoryImpl(db domain.DatabaseInterface, collection string) domain.TaskRepositoryInteface {

	return &taskRepositoryImpl{
		database:   db,
		collection: collection,
	}
}
