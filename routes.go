/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

package main

import (
	"encoding/json"
	"errors"
)

//RunImpl Runs implementation of Invoke function based on function name
func (hdls *HDLS) RunImpl(function string, args []string) ([]byte, error) {

	hdls.logger.Debugf("RunImpl")
	if len(args) != 1 {
		return nil, errors.New("RunImpl: number of argument is invalid")
	}
	disputeContentJSON := args[0]
	disputeContent := CustomerDispute{}
	err := json.Unmarshal([]byte(disputeContentJSON), &disputeContent)
	if err != nil {
		return nil, err
	}
	hdls.logger.Debugf("disputeContent: ", disputeContent)

	// Handle different functions
	switch function {
	//Dispute management functions
	case "addCustomerDispute":
		return nil, hdls.addNewCustomerDispute(disputeContent)

	case "updateCustomerDispute":
		return nil, hdls.updateCustomerDispute(disputeContent)
	case "updatePISPInformation":
		return nil, hdls.updatePISPInformation(disputeContent)
	case "updateMerchantInformation":
		return nil, hdls.updateMerchantInformation(disputeContent)
	case "sendToBankFromPISP":
		return nil, hdls.sendToBankFromPISP(disputeContent)
	case "updateBankInformation":
		return nil, hdls.updateBankInformation(disputeContent)
	case "proposeResolution":
		return nil, hdls.proposeResolution(disputeContent)
	case "approveResolution":
		return nil, hdls.approveResolution(disputeContent)
	case "rejectResolution":
		return nil, hdls.rejectResolution(disputeContent)
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
