package db

import (
	"context"
	"fmt"
	"todo-app/internal/todo"
	"todo-app/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (db *db) Create(ctx context.Context, todo todo.Todo) (string, error) {
	db.logger.Debug("Create todo")
	result, err := db.collection.InsertOne(ctx, todo)
	if err != nil {
		return "", fmt.Errorf("failed to create todo: %v", err)
	}
	db.logger.Debug("Convert insertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	db.logger.Trace(todo)
	return "", fmt.Errorf("failed to convert objectid to hex. Probably oid: %s", oid)
}

func (db *db) Delete(ctx context.Context, id string) error {
	db.logger.Debug("Delete todo")
	result, err := db.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete todo id = %s, error: %v", id, err)
	}
	db.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}

func (db *db) GetAll(ctx context.Context) (todoList []todo.Todo, err error) {
	db.logger.Debug("Get all todo")
	cursor, err := db.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return todoList, fmt.Errorf("failed to get all todo: %v", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var todo todo.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return todoList, fmt.Errorf("failed to decode todo: %v", err)
		}
		todoList = append(todoList, todo)
	}
	db.logger.Tracef("Matched %d documents", len(todoList))

	return todoList, nil
}

func (db *db) Update(ctx context.Context, id string, todo todo.Todo) error {
	db.logger.Debug("Update todo")
	result, err := db.collection.UpdateOne(ctx, bson.M{"_id": id}, todo)
	if err != nil {
		return fmt.Errorf("failed to update todo id = %s, error: %v", id, err)
	}
	db.logger.Tracef("Matched %d and modified %d documents", result.MatchedCount, result.ModifiedCount)
	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) todo.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
