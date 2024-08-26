package config

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

func ConnectMongoDB(env *Env) *mongo.Client {
    clientOptions := options.Client().ApplyURI(env.MongoDBURL)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Check connection
    if err := client.Ping(context.Background(), nil); err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    return client
}
