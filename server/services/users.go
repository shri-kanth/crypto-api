package services

import (
	"crypto-api/server/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := models.User{vars["id"], "name", "address", "email"}
	respondJSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	respondJSON(w, http.StatusOK, models.IdResponse{"qkjwqkjwe"})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	respondJSON(w, http.StatusOK, models.IdResponse{"qkjwqkjwe"})
}
