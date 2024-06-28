package models

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    CustomerID  string             `bson:"customerId" json:"customerId"`
    OrderItems  []OrderItem        `bson:"orderItems" json:"orderItems"`
    TotalAmount float64            `bson:"totalAmount" json:"totalAmount"`
    CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

type OrderItem struct {
    ProductID string  `bson:"productId" json:"productId"`
    Quantity  int     `bson:"quantity" json:"quantity"`
    Price     float64 `bson:"price" json:"price"`
}

var orderCollection *mongo.Collection

// InitializeMongoDB initializes MongoDB collection reference
func InitializeOrderCollection(database *mongo.Database) {
    orderCollection = database.Collection("Orders")
}

// CreateOrder creates a new order in the database.
func (o *Order) CreateOrder() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    o.CreatedAt = time.Now()

    _, err := orderCollection.InsertOne(ctx, o)
    if err != nil {
        return fmt.Errorf("failed to create order: %v", err)
    }

    fmt.Printf("Created order %+v\n", o)
    return nil
}
