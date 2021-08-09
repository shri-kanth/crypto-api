package services

import (
	"crypto-api/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrency(c *gin.Context) {

	// api_key := config.API_KEY
	// vars := mux.Vars(req)
	cid := c.Param("cid")

	cryptoID := models.Currency{cid, "Bitcoin", "This is Bitcoin Description", "random website", 4023.241}
	respondJSON(c.Writer, http.StatusOK, cryptoID)
}

func GetAllCurrencies(c *gin.Context) {

	var arrayOfCryptos = make([]models.Currency, 0)

	arrayOfCryptos = append(arrayOfCryptos, models.Currency{"BTC", "Bitcoin", "This is Bitcoin Description", "random website", 4023.241})
	arrayOfCryptos = append(arrayOfCryptos, models.Currency{"ETH", "Etherium", "This is Etherium Description", "another random website", 424.241})

	respondJSON(c.Writer, http.StatusOK, arrayOfCryptos)
}

func GetCurrencyTimeline(c *gin.Context) {

	// vars := mux.Vars(req)
	cid := c.Param("cid")
	curr := models.CurrencyTimeline{cid,
		[]models.Timeline{{"2021-08-06T10:00:00Z", 324523.34},
			{"2021-08-06T11:00:00Z", 432.45}}}

	respondJSON(c.Writer, http.StatusOK, curr)
}
