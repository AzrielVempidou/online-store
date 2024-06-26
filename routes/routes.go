package routes

import (
	"online-store/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Product routes
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{category}", controllers.GetProductsByCategory).Methods("GET")

	// Shopping Cart routes
	router.HandleFunc("/cart", controllers.CreateCart).Methods("POST")
	router.HandleFunc("/cart", controllers.GetCartItems).Methods("GET")
	router.HandleFunc("/cart/{productId}", controllers.AddToCart).Methods("POST")
	router.HandleFunc("/cart/{productId}", controllers.DeleteFromCart).Methods("DELETE")

	// Order routes
	router.HandleFunc("/checkout", controllers.Checkout).Methods("POST")

	// Auth routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}

