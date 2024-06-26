package controllers

import (
    "encoding/json"
    "net/http"
    "online-store/models"
    "online-store/utils"

    "github.com/gorilla/mux"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
    var cart models.CartItem
    if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    if err := cart.Create(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error creating cart")
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, cart)
}

func AddItemToCart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID := vars["cartId"]

    var item models.CartItem
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    item.CartID = cartID
    if err := item.AddToCart(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error adding item to cart")
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, item)
}

func GetCartItems(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID := vars["cartId"]

    items, err := models.GetCartItems(cartID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error getting cart items")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, items)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]
	cartID := "default_cart_id" // Ganti dengan cara Anda mendapatkan cartID (misalnya, dari sesi pengguna)

	// Check if the product is available
	_, err := models.GetProductByID(productID)
	if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			return
	}

	// Decode JSON request payload to get the requested quantity
	var req struct {
			Quantity int `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
	}

	// Validate the requested quantity
	if req.Quantity <= 0 {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid quantity")
			return
	}

	// Create a new cart item to be added to the cart
	cartItem := models.CartItem{
			CartID:    cartID,
			ProductID: productID,
			Quantity:  req.Quantity,
	}

	// Add the item to the cart
	if err := cartItem.AddToCart(); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error adding item to cart")
			return
	}

	utils.RespondWithJSON(w, http.StatusCreated, cartItem)
}



func DeleteFromCart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID := vars["cartId"]
    itemID := vars["itemId"]

    if err := models.DeleteCartItem(cartID, itemID); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error deleting cart item")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
