package handlers

import (
	"encoding/json"
	"net/http"
	"registration/models"
	"registration/services"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingUser, _ := services.GetUserByUsername(user.Username)
	if existingUser.Username != "" {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	if err := services.RegisterUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}