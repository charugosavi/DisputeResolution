package main

// Transaction and customer identification structure
type TransactionIdentification struct {
	Id            string `json:"id"`         //@PK
	CustomerId    string `json:"customerId"` //@index
	AccountId     string `json:"accountId"`  //@index
	TransactionId string `json:"transactionid"`
}

type TransactionIdentifications struct {
	Data []TransactionIdentification
}

// Transaction information structure
type TransactionInfo struct {
	Id              string  `json:"id"`            //@PK
	TransactionId   string  `json:"transactionId"` //@index
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	TransactionTime string  `json:"time"`
}

type TransactionInfos struct {
	Data []TransactionInfo
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type Customer struct {
	Id                string           `json:"id"` //@PK
	Name              string           `json:"name"`
	Branch            string           `json:"branch"`
	Terminal          string           `json:"terminal"`
	Cashier           string           `json:"cashier"`
	TransactionInfo   *TransactionInfo `json:"transaction"`
	TransactionInfoId string           //@index
	Receipts          []string         `json:"receipts"`
}
type Customers struct {
	Data []Customer
}

type Bank struct {
	Id                string           `json:"id"` //@PK
	Name              string           `json:"name"`
	Branch            string           `json:"branch"`
	Terminal          string           `json:"terminal"`
	Cashier           string           `json:"cashier"`
	TransactionInfo   *TransactionInfo `json:"transaction"`
	TransactionInfoId string           //@index
	Receipts          []string         `json:"receipts"`
}
type Banks struct {
	Data []Bank
}

type Merchant struct {
	Id                string           `json:"id"` //@PK
	Name              string           `json:"name"`
	Branch            string           `json:"branch"`
	Terminal          string           `json:"terminal"`
	Cashier           string           `json:"cashier"`
	TransactionInfo   *TransactionInfo `json:"transaction"`
	TransactionInfoId string           //@index
	Receipts          []string         `json:"receipts"`
}
type Merchants struct {
	Data []Merchant
}

type PISP struct {
	Id                string           `json:"id"` //@PK
	Name              string           `json:"name"`
	Branch            string           `json:"branch"`
	Terminal          string           `json:"terminal"`
	Cashier           string           `json:"cashier"`
	TransactionInfo   *TransactionInfo `json:"transaction"`
	TransactionInfoId string           //@index
	Receipts          []string         `json:"receipts"`
}

type PISPs struct {
	Data []PISP
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type Resolution struct {
	Id                string           `json:"id"`      //@PK
	Outcome           string           `json:"outcome"` //@index
	Description       string           `json:"description"`
	ResolutionTime    string           `json:"resolutionTime"`
	TransactionInfo   *TransactionInfo `json:"transaction"`
	TransactionInfoId string           //@index
}

type Resolutions struct {
	Data []Resolution
}

// Customer initiated dispute structure
type CustomerDispute struct {
	Id                          string                     `json:"disputeId"` //@PK
	TransactionIdentification   *TransactionIdentification `json:"transaction"`
	TransactionIdentificationId string                     //@index
	DisputeType                 string                     `json:"disputetype"`
	Comments                    string                     `json:"comments"`
	Customer                    *Customer                  `json:"customer"`
	CustomerId                  string                     //@index
	Bank                        *Bank                      `json:"bank"`
	BankId                      string                     //@index
	PISP                        *PISP                      `json:"pisp"`
	PISPId                      string                     //@index
	Merchant                    *Merchant                  `json:"merchant"`
	MerchantId                  string                     //@index
	Status                      string                     `json:"status"` //@index
	CreatedDate                 string                     `json:"created"`
	LastUpdated                 string                     `json:"updated"`
	Resolution                  *Resolution                `json:"resolution"`
	ResolutionId                string                     //@index
}

type CustomerDisputes struct {
	Data []CustomerDispute
}
