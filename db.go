package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var runCollection *mongo.Collection

func connectDB() {

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(AppConfig.MongoURI),
	)

	if err != nil {
		log.Fatal("MongoDB connection failed: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}

	db := client.Database("maze")
	runCollection = db.Collection("runs")

	createIndexes()

	log.Println("Connected to MongoDB")
}

func createIndexes() {
	indexModel := mongo.IndexModel{
		Keys: map[string]int{"time": 1},
		Options: options.Index().SetName("time_index"),
	}

	_, _ = runCollection.Indexes().CreateOne(context.Background(), indexModel)

}