package services

import (
	"crypto-api/server/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// vars := mux.Vars(r)
	id := c.Param("id")

	w := c.Writer

	user := models.User{id, "name", "address", "email"}
	respondJSON(w, http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	r := c.Request
	w := c.Writer

	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	respondJSON(w, http.StatusOK, models.IdResponse{"qkjwqkjwe"})
}

func UpdateUser(c *gin.Context) {
	user := models.User{}

	r := c.Request
	w := c.Writer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	respondJSON(w, http.StatusOK, models.IdResponse{"qkjwqkjwe"})
}
