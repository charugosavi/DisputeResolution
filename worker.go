/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

/* ================================ */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
//	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HDLS struct {
	db     shim.ChaincodeStubInterface
	logger *shim.ChaincodeLogger
}

const ()

type TestResponse struct {
	A string "json:a"
	B string "json:b"
}

func (this *HDLS) createKeyEntTable(name string) error {
	if err := this.db.CreateTable(name, []*shim.ColumnDefinition{
		{Name: "Key", Type: shim.ColumnDefinition_STRING, Key: true},
		{"Entity", shim.ColumnDefinition_STRING, false},
	}); err != nil {
		return err
	}
	return nil
}

func (this *HDLS) Init(function string, args []string) ([]byte, error) {
	this.createSchema()

	//if args[0] == "debug" {
	//}
	
	//Initialize the Approver:Creator Organizations combinations with Secret_Token to be used during chaincode event publish
	eventDistributionListJson := args[0]
	var eventDistributions []EventDistribution
	err := json.Unmarshal([]byte(eventDistributionListJson), &eventDistributions)
	if err != nil {
		return nil, err
	}
	for _, eventDistribution := range eventDistributions {	
		this.putEventDistribution(&eventDistribution)
	}
	//Initialization ends

	return nil, nil
}

func (this *HDLS) loadPreference() error {
	/*
		// Custodian
		pref, err := this.getPreference("Custodian")
		var s string
		if pref == nil || err != nil {
			s = DEFAULT_CUSTODIAN
		} else {
			s = pref.Value
		}
		this.custodian = s
	*/

	return nil
}

// Run callback representing the invocation of a chaincode
func (this *HDLS) Invoke(function string, args []string) ([]byte, error) {
	this.loadPreference()

	this.logger.Debugf("worker Invoke: %s, %d", function, len(args));
	ret, err := this.RunImpl(function, args)
	if err != nil {
		this.logger.Debugf("worker Invoke error: %s", err.Error())
	}
	
	// Error processing
//	argsWithComma := strings.Join(args, ",")
//	status := "OK"
//	message := ""
//	payload := fmt.Sprintf("%s(%s)", function, argsWithComma)
//
//	if err != nil {
//		elements := strings.Split(fmt.Sprintf("%v", err), "|")
//		status = elements[0]
//		if len(elements) >= 2 {
//			message = elements[1]
//		}
//	}
//
//	this.putInvokingStatus(&InvokingStatus{this.db.GetTxID(), status, message, payload})
//
//	return ret, nil
	return ret, err
}

// Query callback representing the query of a chaincode
func (this *HDLS) Query(function string, args []string) ([]byte, error) {
	this.loadPreference()

	this.logger.Debugf("worker Query: %s, %d", function, len(args));
	
	if function == "query" {
		var A string // Entities
		var err error

		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
		}

		A = args[0]

		// Get the state from the ledger
		Avalbytes, err := this.db.GetState(A)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
			return nil, errors.New(jsonResp)
		}

		if Avalbytes == nil {
			jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
			return nil, errors.New(jsonResp)
		}

		jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return Avalbytes, nil
	} else {
		result, err := this.QueryImpl(function, args)
		if result != nil {
			j, _ := json.Marshal(result)
			return j, nil
		} else {
			return nil, err
		}
	}
}
