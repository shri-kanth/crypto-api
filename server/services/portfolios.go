package services

import (
	"crypto-api/server/models"
	"net/http"
)

var myPortfolioList = []models.PortfolioList{
	{
		CurrencyId: 1,
		Holdings:   4,
		Price:      19231.0,
		Gains:      10.5,
	},
	{
		CurrencyId: 2,
		Holdings:   100,
		Price:      580.0,
		Gains:      1000.20,
	},
}

var myPortfolio = models.Portfolio{
	TotalBalance: 1056.99, TotalInvestment: 300.0, PortfolioList: myPortfolioList}

var TimeStampList = []models.PortfolioTimeStamp{
	{TimeStamp: "01-01-2021",
		PortfolioByTimeStamp: myPortfolio},
	{TimeStamp: "01-02-2021",
		PortfolioByTimeStamp: myPortfolio},
}

func GetPortfolio(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, myPortfolio)
}

func GetMyPortfolioValue(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, TimeStampList)
}
