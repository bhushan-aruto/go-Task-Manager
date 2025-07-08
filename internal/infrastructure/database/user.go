package database

import (
	"context"

	"time"

	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDatabase) Create(user *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("user")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (db *MongoDatabase) FindByEmail(email string) (*entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("user")
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
