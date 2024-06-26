package utils

import (
	"context"
	"encoding/json"
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

func InitMongoDB(mongoURI string) {
	var err error
	opts := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to check if the connection is established
	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
}

func InitRedis(addr, password string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func GetFromCache(key string) ([]byte, error) {
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func SaveToCache(key string, data interface{}, expiration time.Duration) error {
	ctx := context.Background()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = RedisClient.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func CloseMongoDB() {
	if MongoClient != nil {
		if err := MongoClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error closing MongoDB connection: %v", err)
		}
	}
}

func CloseRedis() {
	if RedisClient != nil {
		if err := RedisClient.Close(); err != nil {
			log.Fatalf("Error closing Redis connection: %v", err)
		}
	}
}
