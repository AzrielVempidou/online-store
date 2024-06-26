package controllers

import (
    "encoding/json"
    "net/http"
    "online-store/models"
    "online-store/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    if err := user.Create(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    token, err := user.Authenticate()
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
