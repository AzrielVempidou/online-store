package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithMessage(w http.ResponseWriter, code int, message string) {
    response := map[string]string{"message": message}
    RespondWithJSON(w, code, response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
    RespondWithMessage(w, code, message)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func RespondWithSuccess(ctx *gin.Context, code int, payload interface{}) {
    ctx.JSON(code, payload)
}