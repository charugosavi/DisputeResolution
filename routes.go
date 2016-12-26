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

//RunImpl Runs implementation of Invoke function based on function name
func (hdls *HDLS) RunImpl(function string, args []string) ([]byte, error) {

	// Handle different functions
	switch function {
	case "import":
		return nil, hdls.imprtJson(args[0])

	//Dispute management functions
	case "addCustomerDispute":
		return nil, hdls.addCustomerDisputeFunction(args)

	case "updateCustomerDispute":
		return nil, hdls.updateCustomerDisputeFunction(args)
	case "updatePISPAssignToMerchant":
		return nil, hdls.updatePISPAssignToMerchantFunction(args)
	default:
		return nil, errors.New("UNKNOWN_INVOCATION|Received unknown function invocation")
	}
}

// QueryImpl Query callback representing the query of a chaincode
func (hdls *HDLS) QueryImpl(function string, args []string) (interface{}, error) {
	switch function {
	case "getCustomerDispute":
		return hdls.getCustomerDispute(args[0])
	case "listCustomerDisputes":
		return hdls.listCustomerDisputes()

	default:
		return nil, errors.New("Received unknown function query")
	}

}
