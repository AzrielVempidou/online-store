package controllers

import (
    "encoding/json"
    "net/http"
    "online-store/models"
    "online-store/utils"

    "github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := models.GetAllProducts()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error getting products")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, products)
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]

    products, err := models.GetProductsByCategory(category)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error getting products by category")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, products)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        RespondWithError(w, http.StatusInternalServerError, "Error converting to JSON")
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(jsonData)
}

func RespondWithError(w http.ResponseWriter, i int, s string) {
	panic("unimplemented")
}
