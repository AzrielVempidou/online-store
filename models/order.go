package models

import (
    "context"
    "online-store/utils"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    CustomerID  string             `bson:"customerId" json:"customerId"`
    TotalAmount float64            `bson:"totalAmount" json:"totalAmount"`
    Status      string             `bson:"status" json:"status"`
    CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
    Items       []OrderItem        `bson:"items" json:"items"`
}

type OrderItem struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    OrderID   primitive.ObjectID `bson:"orderId,omitempty" json:"orderId,omitempty"`
    ProductID primitive.ObjectID `bson:"productId,omitempty" json:"productId,omitempty"`
    Quantity  int                `bson:"quantity,omitempty" json:"quantity,omitempty"`
    Price     float64            `bson:"price,omitempty" json:"price,omitempty"`
    CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (o *Order) Create() error {
    collection := utils.MongoClient.Database("store").Collection("orders")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    o.CreatedAt = time.Now()
    _, err := collection.InsertOne(ctx, o)
    return err
}

func (oi *OrderItem) MarshalBSON() ([]byte, error) {
    oi.CreatedAt = time.Now()
    return bson.Marshal(oi)
}

func (oi *OrderItem) UnmarshalBSON(data []byte) error {
    type Alias OrderItem
    var temp struct {
        Alias
        CreatedAt *time.Time `bson:"createdAt"`
    }
    if err := bson.Unmarshal(data, &temp); err != nil {
        return err
    }
    *oi = OrderItem(temp.Alias)
    if temp.CreatedAt != nil {
        oi.CreatedAt = *temp.CreatedAt
    }
    return nil
}
