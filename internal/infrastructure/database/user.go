package database

import (
	"context"

	"time"

	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"github.com/bhushan-aruto/go-task-manager/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepo struct {
	db *MongoDatabase
}

func NewUserRepo(db *MongoDatabase) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := r.db.Collection("user")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *userRepo) FindByEmail(email string) (*entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := r.db.Collection("user")
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
