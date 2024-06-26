package main

import (
    "log"
    "net/http"
    "online-store/routes"
    "online-store/utils"

    "github.com/gorilla/mux"
)

func main() {
		// Inisialisasi koneksi MongoDB
		utils.InitMongoDB("mongodb://localhost:27017")

		// Inisialisasi koneksi Redis
		utils.InitRedis("localhost:6379", "")
    router := mux.NewRouter()
    routes.RegisterRoutes(router)

    log.Fatal(http.ListenAndServe(":8000", router))
}
