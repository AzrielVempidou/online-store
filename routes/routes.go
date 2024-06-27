package routes

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "online-store/controllers" // Sesuaikan dengan path yang tepat
)

// TestHandler is a simple handler for testing purposes
func TestHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Example app listening on port %s", "8000")
}

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

    // Test route
    router.HandleFunc("/", TestHandler).Methods("GET")
}
