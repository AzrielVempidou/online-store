package models

import (
    "context"
    "errors"
    "online-store/utils"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name      string             `bson:"name" json:"name"`
    Email     string             `bson:"email" json:"email"`
    Password  string             `bson:"password" json:"-"`
    CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func (u *User) Create() error {
    collection := utils.MongoClient.Database("store").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    u.CreatedAt = time.Now()
    _, err = collection.InsertOne(ctx, u)
    return err
}

func (u *User) Authenticate() (string, error) {
    collection := utils.MongoClient.Database("store").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user User
    err := collection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&user)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: u.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
