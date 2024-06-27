package controllers

import (
    "encoding/json"
    "log"
    "net/http"
    "online-store/models"
    "online-store/utils"
    "time"

    "github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := models.GetAllProducts()
    if err != nil {
        log.Printf("Error retrieving products: %v", err)  // Log error detail
        RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    RespondWithJSON(w, http.StatusOK, products)
}


func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]

    // Check if the category data exists in Redis
    cacheKey := "products:" + category
    cachedData, err := utils.GetFromCache(cacheKey)
    if err == nil && cachedData != nil {
        // If found in Redis cache, respond with cached data
        RespondWithJSON(w, http.StatusOK, cachedData)
        return
    }

    // If not found in Redis cache, fetch from MongoDB
    products, err := models.GetProductsByCategory(category)
    if err != nil {
        RespondWithError(w, http.StatusInternalServerError, "Error getting products by category")
        return
    }

    // Save fetched data to Redis for future requests
    expiration := 10 * time.Minute // Example: cache for 10 minutes
    if err := utils.SaveToCache(cacheKey, products, expiration); err != nil {
        log.Printf("Failed to save data to Redis: %v", err)
    }

    RespondWithJSON(w, http.StatusOK, products)
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

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    errorResponse := map[string]string{"error": message}
    jsonData, err := json.Marshal(errorResponse)
    if err != nil {
        // If there's an error encoding the error response, log it
        log.Printf("Error encoding error response: %v", err)
        return
    }
    w.Write(jsonData)
}
