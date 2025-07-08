package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	client   *mongo.Client
	database *mongo.Database
}

func Newdatabase(databaseurl string ,config *con) *MongoDatabase {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseurl))

	if err != nil {
		log.Fatalf("error occured while connecting mongo db")

	}




}
