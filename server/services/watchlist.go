package services

import (
	"crypto-api/server/models"
	"encoding/json"
	"net/http"
)

var watchlist = []models.CryptoCurrency{
	{CryptoCurrencyId: "BTC"},
	{CryptoCurrencyId: "ETH"},
}

func GetWatchList(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, watchlist)
}

func UpdateWatchList(w http.ResponseWriter, r *http.Request) {

	var newCryptoCurrency models.CryptoCurrency
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCryptoCurrency); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	watchlist = append(watchlist, newCryptoCurrency)

	defer r.Body.Close()
	respondJSON(w, http.StatusOK, watchlist)
}
