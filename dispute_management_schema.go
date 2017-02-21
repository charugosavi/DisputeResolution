package main

// Transaction and customer identification structure
type TransactionIdentification struct {
	Id            string `json:"id"` //@PK
	CustomerId    string `json:"customerId"`
	AccountId     string `json:"accountId"`
	TransactionId string `json:"transactionid"`
}

// Transaction information structure
type TransactionInfo struct {
	Id              string `json:"id"` //@PK
	TransactionId   string `json:"transactionId"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	TransactionTime string `json:"time"`
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type Customer struct {
	Id              string           `json:"id"` //@PK
	CustomerId      string           `json:"customerId"`
	AccountId       string           `json:"accountId"`
	Name            string           `json:"name"`
	Comments        string           `json:"comments"`
	TransactionInfo *TransactionInfo `json:"transaction"`
	Receipts        []string         `json:"receipts"`
}

type Bank struct {
	Id              string           `json:"id"` //@PK
	Name            string           `json:"name"`
	Branch          string           `json:"branch"`
	Terminal        string           `json:"terminal"`
	Cashier         string           `json:"cashier"`
	Comments        string           `json:"comments"`
	TransactionInfo *TransactionInfo `json:"transaction"`
	Receipts        []string         `json:"receipts"`
}

type Merchant struct {
	Id              string           `json:"id"` //@PK
	Name            string           `json:"name"`
	Branch          string           `json:"branch"`
	Terminal        string           `json:"terminal"`
	Cashier         string           `json:"cashier"`
	Comments        string           `json:"comments"`
	TransactionInfo *TransactionInfo `json:"transaction"`
	Receipts        []string         `json:"receipts"`
}

type PISP struct {
	Id              string           `json:"id"` //@PK
	Name            string           `json:"name"`
	Branch          string           `json:"branch"`
	Terminal        string           `json:"terminal"`
	Cashier         string           `json:"cashier"`
	Comments        string           `json:"comments"`
	TransactionInfo *TransactionInfo `json:"transaction"`
	Receipts        []string         `json:"receipts"`
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type Resolution struct {
	Id              string           `json:"id"` //@PK
	Type            string           `json:"type"`
	Outcome         string           `json:"outcome"`
	Description     string           `json:"description"`
	ResolutionTime  string           `json:"resolutionTime"`
	TransactionInfo *TransactionInfo `json:"transaction"`
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type ResolutionExecution struct {
	Id              string           `json:"id"` //@PK
	Owner           string           `json:"owner"`
	TransactionInfo *TransactionInfo `json:"transaction"`
	Comments        string           `json:"comments"`
	Receipts        []string         `json:"receipts"`
}

// Customer initiated dispute structure
type CustomerDispute struct {
	Id                  string           `json:"disputeId"` //@PK
	TransactionInfo     *TransactionInfo `json:"transaction"`
	DisputeType         string           `json:"disputetype"`
	Customer            *Customer        `json:"customer"`
	Bank                *Bank            `json:"bank"`
	PISP                *PISP            `json:"pisp"`
	Merchant            *Merchant        `json:"merchant"`
	Status              string           `json:"status"`
	CreatedDate         string           `json:"created"`
	LastUpdated         string           `json:"updated"`
	Resolution          *Resolution      `json:"resolution"`
	ResolutionExecution []*Resolution    `json:"execution"`
	Owner               []string         `json:"owner"`
	Audit               []string         `json:"audit"`
}
