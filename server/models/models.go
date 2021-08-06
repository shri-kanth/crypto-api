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

type Currency struct {
	Id          string
	Name        string
	Description string
	Logo_url    string
	Market_cap  float64
}

type Timeline struct {
	Timestamp      string
	Currency_value float64
}

type CurrencyTimeline struct {
	Id        string
	Timelines []Timeline
}
