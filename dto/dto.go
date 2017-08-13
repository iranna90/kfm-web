package dto

import "time"

type Person struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type DailyMilkTransaction struct {
	NumberOfLiters  int8 `json:"numberOfLiters"`
	TotalPriceOfDay int `json:"totalPriceOfTheDay"`
	Balance         int64 `json:"balance"`
	Day             time.Time `json:"time"`
	PersonName      string `json:"personName"`
}

type Payment struct {
	Amount          int64 `json:"amount"`
	PaidTo          string `json:"paidTo"`
	Day             time.Time `json:"day"`
	RemainingAmount int64 `json:"remainingBalance"`
}

type Balance struct {
	Amount    int64
	Modified  time.Time
	PersonRef int64
}

type Address struct {
	PhoneNumber int64
	FullAddress string
	PersonRef   int64
}
