package services

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"crypto-api/server/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := models.User{vars["id"],"saketh","lingampalli","saketh.kokonda@zemosolabs.com", "91728392039", "random_url"}
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}