package database

import (
	"context"
	"errors"
	"time"

	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"github.com/bhushan-aruto/go-task-manager/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepo struct {
	db *MongoDatabase
}

func NewTaskRepo(database *MongoDatabase) repository.TaskRepository {
	return &taskRepo{
		db: database,
	}
}

func (r *taskRepo) Creat(task *entity.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := r.db.Collection("tasks")

	_, err := collection.InsertOne(ctx, task)

	return err
}

func (r *taskRepo) GetTaskById(id string) (*entity.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// var task entity.Task
	task := new(entity.Task)
	collection := r.db.Collection("tasks")

	if err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepo) UpdateTask(task *entity.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := r.db.Collection("tasks")

	if task.ID.IsZero() {
		return errors.New("invalid task ID")
	}
	filter := bson.M{"_id": task.ID}
	update := bson.M{"$set": bson.M{
		"title":       task.Title,
		"description": task.Description,
		"Completed":   task.Completed,
		"updated_at":  time.Now(),
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *taskRepo) DeleteTaskById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.db.Collection("tasks")

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objId})
	return err

}

func (r *taskRepo) ListByUser(userId string) ([]*entity.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := r.db.Collection("tasks")

	cursor, err := collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*entity.Task
	for cursor.Next(ctx) {
		task := new(entity.Task)
		if err := cursor.Decode(task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
