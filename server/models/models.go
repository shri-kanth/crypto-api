package models

type User struct {
	Id string
	Name string
	Address string
	Email string
	Phone string
	Picture string
}

type IdResponse struct {
	Id string
}