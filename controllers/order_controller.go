package controllers

import (
    "encoding/json"
    "net/http"
    "online-store/models"
    "online-store/utils"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    if err := order.Create(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error creating order")
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, order)
}
