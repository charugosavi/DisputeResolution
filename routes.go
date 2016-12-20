/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

package main

import "errors"

type Status struct {
	code string
	msg  string
}

type Response struct {
	Contents interface{}
}

func (this *HDLS) RunImpl(function string, args []string) ([]byte, error) {

	// Handle different functions
	switch function {
	case "import":
		return nil, this.imprtJson(args[0])

	//Dispute management functions
	case "addCustomerDispute":
		return nil, this.addCustomerDisputeFunction(args)

	default:
		return nil, errors.New("UNKNOWN_INVOCATION|Received unknown function invocation")
	}
}

// Query callback representing the query of a chaincode
func (this *HDLS) QueryImpl(function string, args []string) (interface{}, error) {
	switch function {
	case "getCustomerDispute":
		return this.getCustomerDispute(args[0])
	case "listCustomerDisputes":
		return this.listCustomerDisputes()

	default:
		return nil, errors.New("Received unknown function query")
	}

}
