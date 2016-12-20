package main

// Transaction and customer identification structure
type TransactionIdentification struct {
	Id            string `json:"id"`         //@PK
	CustomerId    string `json:"customerId"` //@index
	AccountId     string `json:"accountId"`  //@index
	TransactionId string `json:"transactionid"`
}

// Transaction information structure
type TransactionInfo struct {
	Id              string  `json:"id"` //@PK
	TransactionId   string  `json:"transactionId"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	TransactionTime string  `json:"time"`
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type InvolvedParty struct {
	Id          string           `json:"id"` //@PK
	Name        string           `json:"name"`
	Branch      string           `json:"branch"`
	Terminal    string           `json:"terminal"`
	Cashier     string           `json:"cashier"`
	Transaction *TransactionInfo `json:"transaction"`
	Receipts    []string         `json:"receipts"`
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type Resolution struct {
	Id             string           `json:"id"`      //@PK
	Outcome        string           `json:"outcome"` //@index
	Description    string           `json:"description"`
	ResolutionTime string           `json:"resolutionTime"`
	Transaction    *TransactionInfo `json:"transaction"`
}

// Customer initiated dispute structure
type CustomerDispute struct {
	DisputeId   string                     `json:"disputeId"` //@PK
	Transaction *TransactionIdentification `json:"transaction"`
	DisputeType string                     `json:"disputetype"`
	Comments    string                     `json:"comments"`
	Customer    *InvolvedParty             `json:"customer"`
	Bank        *InvolvedParty             `json:"bank"`
	PISP        *InvolvedParty             `json:"pisp"`
	Merchant    *InvolvedParty             `json:"merchant"`
	Status      string                     `json:"status"` //@index
	CreatedDate string                     `json:"created"`
	LastUpdated string                     `json:"updated"`
	Resolution  *Resolution                `json:"resolution"`
}
