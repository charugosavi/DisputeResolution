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
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HDLSMaster struct {
	/* Do not include any variables here */
}

func (this *HDLSMaster) newWorker(db shim.ChaincodeStubInterface) *HDLS {
	logger := shim.NewLogger("HDLS")
	return &HDLS{db: db, logger: logger}
}

func (this *HDLSMaster) Init(db shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	worker := this.newWorker(db)
	return worker.Init(function, args)
}

// Run callback representing the invocation of a chaincode
func (this *HDLSMaster) Invoke(db shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	worker := this.newWorker(db)
	return worker.Invoke(function, args)
}

// Query callback representing the query of a chaincode
func (this *HDLSMaster) Query(db shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	worker := this.newWorker(db)
	return worker.Query(function, args)
}

func main() {
	err := shim.Start(new(HDLSMaster))
	if err != nil {
		fmt.Printf("Error starting HDLS Master: %s", err)
	}
}
