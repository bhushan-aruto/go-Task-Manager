package database

import (
	"context"
	"log"
	"time"

	"github.com/bhushan-aruto/go-task-manager/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	client   *mongo.Client
	database *mongo.Database
}

func ConnecttoMOngo(cfg *config.Config) *MongoDatabase {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DatabaseUrl))

	if err != nil {
		log.Fatalf("error occured while connecting mongo db")

	}

	db := mongoClient.Database(cfg.DatabaseName)

	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalln("error occured while connecting the  mongoDB ", err.Error())
	}

	log.Println("Connected to the mongoDB")

	return &MongoDatabase{
		client:   mongoClient,
		database: db,
	}
}

func (db *MongoDatabase) Collection(name string) *mongo.Collection {
	return db.database.Collection(name)

}
