package services

import (
	"crypto-api/server/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCurrency(res http.ResponseWriter, req *http.Request) {

	// api_key := config.API_KEY
	vars := mux.Vars(req)

	cryptoID := models.Currency{vars["cid"], "Bitcoin", "This is Bitcoin Description", "random website", 4023.241}
	respondJSON(res, http.StatusOK, cryptoID)
}

func GetAllCurrencies(res http.ResponseWriter, req *http.Request) {

	var arrayOfCryptos = make([]models.Currency, 0)

	arrayOfCryptos = append(arrayOfCryptos, models.Currency{"BTC", "Bitcoin", "This is Bitcoin Description", "random website", 4023.241})
	arrayOfCryptos = append(arrayOfCryptos, models.Currency{"ETH", "Etherium", "This is Etherium Description", "another random website", 424.241})

	respondJSON(res, http.StatusOK, arrayOfCryptos)
}

func GetCurrencyTimeline(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	curr := models.CurrencyTimeline{vars["cid"],
		[]models.Timeline{{"2021-08-06T10:00:00Z", 324523.34},
			{"2021-08-06T11:00:00Z", 432.45}}}

	respondJSON(res, http.StatusOK, curr)
}
