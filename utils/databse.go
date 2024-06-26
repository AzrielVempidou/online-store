package utils

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/go-redis/redis/v8"
)

var (
    MongoClient *mongo.Client
    RedisClient *redis.Client
)

func InitMongoDB(uri string) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    MongoClient = client
}

func InitRedis(addr, password string) {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       0,
    })
    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
}
