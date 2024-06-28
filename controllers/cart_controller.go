package controllers

import (
	"encoding/json"
	"net/http"
	"online-store/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type CartController struct{}

func NewCartController() *CartController {
    return &CartController{}
}

// AddToCart adds an item to the cart.
func (c *CartController) AddToCart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID := vars["cartId"]
    productID := vars["productId"]

    var input struct {
        Quantity int `json:"quantity"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    cartItem := models.CartItem{
        CartID:    cartID,
        ProductID: productID,
        Quantity:  input.Quantity,
    }

    if err := cartItem.AddToCart(); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, gin.H{"message": "Product added to cart"})
}

// GetCartItems retrieves all items in the cart.
func (c *CartController) GetCartItems(w http.ResponseWriter, r *http.Request) {
    cartID := mux.Vars(r)["cartId"]

    items, err := models.GetCartItems(cartID)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, items)
}

// DeleteFromCart deletes an item from the cart.
func (c *CartController) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
    cartID := mux.Vars(r)["cartId"]
    itemID := mux.Vars(r)["itemId"]

    if err := models.DeleteCartItem(cartID, itemID); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, gin.H{"message": "Item deleted from cart"})
}

// Utility function to respond with JSON data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

// Utility function to respond with error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
    respondWithJSON(w, code, gin.H{"error": msg})
}
