package models

type Transaction struct {
	Id          string
	CurrencyId  string
	Units       float64
	TotalValue  float64
	Status      string
	TimeStamp   string
	PaymentType string
	PaymentMode string
}

type IdResponse struct {
	Id string
}

type ErrorResponse struct {
	ErrorMessage string
}
