// models/cart_item.go

package models

import (
    "context"
    "errors"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type CartItem struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    CartID    string             `bson:"cartId" json:"cartId"`
    ProductID string             `bson:"productId" json:"productId"`
    Quantity  int                `bson:"quantity" json:"quantity"`
    CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

var cartCollection *mongo.Collection

// InitializeMongoDB initializes MongoDB collection reference
func InitializeMongoDB(database *mongo.Database) {
    cartCollection = database.Collection("carts")
}

// Create creates a new shopping cart.
func (c *CartItem) Create() error {
    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set created timestamp
    c.CreatedAt = time.Now()

    // Insert document
    _, err := cartCollection.InsertOne(ctx, c)
    if err != nil {
        return fmt.Errorf("failed to create cart: %v", err)
    }

    fmt.Printf("Created cart %+v\n", c)
    return nil
}

// AddToCart adds the cart item to the specified cart in the database.
func (ci *CartItem) AddToCart() error {
    // Validate input
    if ci.CartID == "" {
        return errors.New("cartID is required")
    }
    if ci.ProductID == "" {
        return errors.New("productID is required")
    }
    if ci.Quantity <= 0 {
        return errors.New("quantity must be greater than 0")
    }

    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set created timestamp
    ci.CreatedAt = time.Now()

    // Insert document
    _, err := cartCollection.InsertOne(ctx, ci)
    if err != nil {
        return fmt.Errorf("failed to insert item into cart: %v", err)
    }

    fmt.Printf("Added item %+v to cart %s\n", ci, ci.CartID)
    return nil
}

// GetCartItems retrieves all cart items associated with a cartID.
func GetCartItems(cartID string) ([]CartItem, error) {
    var items []CartItem
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := cartCollection.Find(ctx, bson.M{"cartId": cartID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var item CartItem
        if err = cursor.Decode(&item); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

// DeleteCartItem deletes a specific cart item from the database.
func DeleteCartItem(cartID, itemID string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    itemObjectID, err := primitive.ObjectIDFromHex(itemID)
    if err != nil {
        return err
    }
    _, err = cartCollection.DeleteOne(ctx, bson.M{"_id": itemObjectID, "cartId": cartID})
    return err
}
