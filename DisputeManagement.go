package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Transaction and customer identification structure
type TransactionIdentification struct {
	CustomerId string `json:"customerId"`
	AccountId string `json:"accountId"`
	TransactionId string `json:"transactionid"`
}

// Transaction information structure
type TransactionInfo struct {
	TransactionId string `json:"transactionId"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	TransactionTime string `json:"time"`
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type InvolvedParty struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Branch string `json:"branch"`
	Terminal string `json:"terminal"`
	Cashier string `json:"cashier"`
	Transaction TransactionInfo `json:"transaction"`
	Receipts []string `json:"receipts"`	
}

// Invovled party information structure. Used to represent Merchant, PISP and Bank transaction information.
type Resolution struct {
	Id string `json:"id"`
	Outcome string `json:"outcome"`
	Description string `json:"description"`
	ResolutionTime string `json:"resolutionTime"`
	Transaction TransactionInfo `json:"transaction"` 	
}

// Customer initiated dispute structure
type CustomerDispute struct {
	DisputeId string `json:"disputeId"`
	Transaction TransactionIdentification `json:"transaction"`
	DisputeType string `json:"disputetype"`
	Comments string `json:"comments"`
	Customer InvolvedParty `json:"customer"`
	Bank InvolvedParty `json:"bank"`
	PISP InvolvedParty `json:"pisp"`
	Merchant InvolvedParty `json:"merchant"`	
	Status string `json:"status"`
	CreatedDate string `json:"created"`
	LastUpdated string `json:"updated"`
	Resolution Resolution `json:"resolution"`	
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Dispute Management chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Println("Init is running ")
	
	return nil, nil
}

// Invoke is an entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Println("invoke is running " + function)

	if function == "create" {
		var dispute CustomerDispute
		err := json.Unmarshal([]byte(args[0]), &dispute)
		//timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		dispute.DisputeId = RandStringBytesRmndr(10)

		disputeRecordJSON, err := json.Marshal(dispute)
		
		if(err != nil){
			return nil, err		
		}
		
		// store the JSON on ledger
		err = stub.PutState(dispute.DisputeId, disputeRecordJSON) //write the variable into the chaincode state
		if err != nil {
			return nil, err
		}
	} else if function == "update" {
		var disputeId string
		disputeId = args[0]
		var updatedDispute CustomerDispute
		err := json.Unmarshal([]byte(args[1]), &updatedDispute)

		disputeRecordJSON, err := json.Marshal(updatedDispute)
		
		if(err != nil){
			return nil, err		
		}
		
		// store the JSON on ledger
		err = stub.PutState(disputeId, disputeRecordJSON) //write the variable into the chaincode state
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// Query is an entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Println("query is running " + function)
	
	if function == "read" {
	
		var key, jsonResp string
		var err error

		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
		}

		key = args[0]
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		return valAsbytes, nil
	}
	return nil, nil
}

func RandStringBytesRmndr(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}