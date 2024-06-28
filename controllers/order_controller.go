package controllers

import (
	"encoding/json"
	"net/http"
	"online-store/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func NewOrderController() *OrderController {
    return &OrderController{}
}

// Checkout processes the checkout and creates an order.
func (oc *OrderController) Checkout(w http.ResponseWriter, r *http.Request) {
    var order models.Order

    // Parse JSON request body
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Validate order (example: calculate total amount)
    order.TotalAmount = calculateTotalAmount(order.OrderItems)

    // Create order in the database
    if err := order.CreateOrder(); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    // Return success message
    respondWithJSON(w, http.StatusOK, gin.H{"message": "Order created successfully"})
}

// Utility function to calculate total amount based on order items
func calculateTotalAmount(items []models.OrderItem) float64 {
    var total float64
    for _, item := range items {
        total += float64(item.Quantity) * item.Price
    }
    return total
}
