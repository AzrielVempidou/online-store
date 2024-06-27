package controllers

import (
	"log"
	"net/http"
	"online-store/models"
	"online-store/utils"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := models.GetAllProducts()
    if err != nil {
        log.Printf("Error retrieving products: %v", err)  // Log error detail
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, products)
}


func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]

    // Lakukan sesuatu dengan kategori, misalnya query ke database
    // Contoh sederhana untuk merespons dengan kategori yang diterima
    response := "Menampilkan produk untuk kategori: " + category
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(response))
}

