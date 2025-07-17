package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task_manager/db"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllTasks is a function that return
// all tasks currently available
func GetAllTasks() ([]models.Task, error) {
	tasks := db.Client.Database(db.DBName).Collection("tasks")

	// set timeout for the function
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := tasks.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []models.Task

	for cursor.TryNext(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Printf("Failed to decode task: %s", err.Error())
			continue
		}
		results = append(results, task)
	}
	return results, nil
}

// AddTask is a method used to add
// a new task into the tasks slice
func AddTask(task models.Task) (models.Task, error) {
	tasks := db.Client.Database(db.DBName).Collection("tasks")

	// set timeout for insert operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := tasks.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, err
	}

	// set the generated ID back on the task
	task.ID = result.InsertedID.(primitive.ObjectID)

	return task, nil
}

// GetTaskByID iterates over all tasks and returns the task
// if not found, returns an error
func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	tasks := db.Client.Database(db.DBName).Collection("tasks")
	// set timeout for fetch operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// prepare filter
	filter := bson.D{{Key: "_id", Value: id}}
	// find the document
	var task models.Task
	err := tasks.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Task{}, fmt.Errorf("task does not exist: %w", err)
		}
		return models.Task{}, err
	}
	return task, nil
}

// UpdateTask finds a task by id and updates it with the new modesl
// if task is not found it returns an error
func UpdateTask(id primitive.ObjectID, modTask models.Task) error {
	tasks := db.Client.Database(db.DBName).Collection("tasks")
	// set timeout for fetch operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// prepare filter data
	filter := bson.D{{Key: "_id", Value: id}}
	// prepare the update data
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: modTask.Title},
			{Key: "description", Value: modTask.Description},
			{Key: "due_date", Value: modTask.DueDate},
			{Key: "status", Value: modTask.Status},
		},
		},
	}

	result, err := tasks.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	// if document was not found
	if result.MatchedCount == 0 {
		return fmt.Errorf("task not found")
	}
	return nil
}

// DeleteTaskByID removes a task from tasks slice
// returns an error if task is not found
func DeleteTaskByID(id primitive.ObjectID) error {
	tasks := db.Client.Database(db.DBName).Collection("tasks")
	// set timeout for fetch operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// prepare the filter
	filter := bson.D{{Key: "_id", Value: id}}

	result := tasks.FindOneAndDelete(ctx, filter)
	// check if the deletion was a success
	var deletedTask models.Task
	err := result.Decode(&deletedTask)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("task with id %s is not found", id)

		}
		return err
	}

	return nil
}
