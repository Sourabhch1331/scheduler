package database

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"scheduler-server/internal"
	"time"
)

type MongoDb struct {
	client *mongo.Client
}

func NewMongoDb() *MongoDb {
	return &MongoDb{}
}

func (db *MongoDb) Open() error {
	uri := internal.GetConfigString("mongo.connection_string")
	if uri == "" {
		return errors.New("connection string not found in config")
	}

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB connection test failed: %v", err)
	}

	db.client = client
	fmt.Println("Connected to MongoDB")
	return nil
}

func (db *MongoDb) Close() {
	if db == nil || db.client == nil {
		return
	}

	if err := db.client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect MongoDB: %v", err)
	}
}
