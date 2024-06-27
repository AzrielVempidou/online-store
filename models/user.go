package models

import (
	"context"
	"errors"
	"online-store/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

// Create inserts a new user into the database
func (u *User) Create() error {
	collection := utils.MongoClient.Database("store").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if email already exists
	var existingUser User
	err := collection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&existingUser)
	if err == nil {
		// If email already exists, return an error
		return errors.New("email already exists")
	} else if err != nil && err != mongo.ErrNoDocuments {
		// If an unexpected error occurs, return the error
		return err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Set CreatedAt timestamp
	u.CreatedAt = time.Now()

	// Insert the user into the database
	_, err = collection.InsertOne(ctx, u)

	return nil
}

// Authenticate checks the user credentials and generates a JWT token
func (u *User) Authenticate() (string, error) {
	collection := utils.MongoClient.Database("store").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	tokenString, err := generateJWTToken(u.Email)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// generateJWTToken generates a JWT token with user email claims
func generateJWTToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
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
