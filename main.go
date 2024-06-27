package main

import (
	"log"
	"net/http"
	"online-store/routes"
	"online-store/utils"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Inisialisasi koneksi MongoDB
	// Memuat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}
	utils.InitMongoDB(mongoURI)

	// Inisialisasi koneksi Redis
	// utils.InitRedis("localhost:6379", "")

	// Membuat router baru
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Mendapatkan port dari variabel lingkungan atau menggunakan default port 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Menampilkan pesan ke log bahwa aplikasi sudah berjalan
	log.Printf("Example app listening on port %s", port)

	// Memulai server
	log.Fatal(http.ListenAndServe(":"+port, router))
}