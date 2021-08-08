package models

type User struct {
	Id      string
	Name    string
	Address string
	Email   string
}

type IdResponse struct {
	Id string
}

//------------------------ Portfolio Structures -------------------------------------

type Portfolio struct {
	TotalBalance    float64         `json:"totalBalnce"`
	TotalInvestment float64         `json:"total Investment"`
	PortfolioList   []PortfolioList `json:"PortfolioList"`
}

type PortfolioList struct {
	CurrencyId int     `json:"currencyId"`
	Holdings   int     `json:"holdings"`
	Price      float64 `json:"price"`
	Gains      float64 `json:"gains"`
}

type PortfolioTimeStamp struct {
	TimeStamp            string    `json:"timeStamp"`
	PortfolioByTimeStamp Portfolio `json:"portfolio"`
}

// ----------------------------- WatchList--------------------------------

type CryptoCurrency struct {
	CryptoCurrencyId string `json:"cryptoCurrencyId"`
}
