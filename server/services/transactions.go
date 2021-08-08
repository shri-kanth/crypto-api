package services

import (
	"crypto-api/server/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var transactions = []models.Transaction{
	{"1", "BTC", 0.00224047, 100, "Success", "2021-08-08T15:13:22.801Z", "Buy", "UPI"},
	{"2", "ETH", 0.1721571, 530.56, "Success", "2021-08-08T15:13:22.801Z", "Buy", "UPI"},
	{"3", "LTC", 0.098044, 14.82, "Success", "2021-08-08T15:13:22.801Z", "Sell", "UPI"},
	{"4", "XRP", 24.900243, 20, "Pending", "2021-08-08T15:13:22.801Z", "Buy", "UPI"},
	{"5", "ADA", 135.6871, 200, "Failure", "2021-08-08T15:13:22.801Z", "Buy", "UPI"},
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 0)
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 0)

	// for i := 1; int64(i) <= limit && offset > 0; i++ {
	// 	transactions = append(transactions, models.Transaction{strconv.Itoa(i), "BTC", 0.0023, 100, "Success", "2021-08-08T15:13:22.801Z", "Buy", "UPI"})
	// }

	flag := false

	if int(limit) <= 0 || int(limit) > len(transactions) {
		flag = true
		respondJSON(w, http.StatusBadRequest, models.ErrorResponse{"Limit value is invalid"})
	}

	if int(offset) <= 0 {
		respondJSON(w, http.StatusBadRequest, models.ErrorResponse{"Offset value is invalid"})
		return
	}

	if flag {
		return
	}

	respondJSON(w, http.StatusOK, transactions[len(transactions)-int(limit):])
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	transactionId := vars["transactionId"]
	id, _ := strconv.ParseInt(transactionId, 10, 0)

	if userId == "nil" {
		respondJSON(w, http.StatusBadRequest, models.ErrorResponse{"Invalid User ID"})
		return
	}

	if id <= 0 || int(id) > len(transactions) {
		respondJSON(w, http.StatusBadRequest, models.ErrorResponse{"Invalid transaction ID"})
		return
	}
	respondJSON(w, http.StatusOK, transactions[int(id)-1])
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaction); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	respondJSON(w, http.StatusOK, transaction)
}
