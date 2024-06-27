package models

import (
    "context"
    "errors"
    "online-store/utils"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string             `bson:"name" json:"name"`
    Description string             `bson:"description" json:"description"`
    Category    string             `bson:"category" json:"category"`
    Price       float64            `bson:"price" json:"price"`
    Stock       int                `bson:"stock" json:"stock"`
    CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

var productCollection *mongo.Collection

// EnsureCollection initializes the MongoDB collection for products
func EnsureCollection() {
    if productCollection == nil {
        productCollection = utils.MongoClient.Database("store").Collection("products")
    }
}

// GetProductByID retrieves a product by its ID
func GetProductByID(id string) (*Product, error) {
    EnsureCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Convert id to MongoDB ObjectID
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    // Search for product with given ObjectID
    var product Product
    err = productCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, errors.New("product not found")
        }
        return nil, err
    }

    return &product, nil
}

// GetAllProducts retrieves all products from the collection
func GetAllProducts() ([]Product, error) {
    EnsureCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := productCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var products []Product
    for cursor.Next(ctx) {
        var product Product
        if err := cursor.Decode(&product); err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    return products, nil
}

// GetProductsByCategory retrieves products by category
func GetProductsByCategory(category string) ([]Product, error) {
    EnsureCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := productCollection.Find(ctx, bson.M{"category": category})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var products []Product
    for cursor.Next(ctx) {
        var product Product
        if err = cursor.Decode(&product); err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}
