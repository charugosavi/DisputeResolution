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
	"errors"

	"github.com/hyperledger/fabric/core/crypto/primitives"
)

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
	case "deleteAllDocument":
		return nil, this.deleteAllDocument()

	//Dispute management functions
	case "addCustomerDispute":
		return nil, this.addCustomerDispute(args)
	case "deleteCustomerDispute":
		return nil, this.deleteCustomerDispute(args)
	case "updateCustomerDispute":
		return nil, this.overwriteCustomerDispute(args)
	case "approveSoW":
		return nil, this.approveSoWFunction(args)
	case "createBaseDocumentAttachment":
		return nil, this.createBaseDocumentAttachmentFunction(args)
	case "createAttachmentToDoc":
		return nil, this.createAttachmentToDocFunction(args)

	//MSA Functions
	case "createMSA":
		return nil, this.createMSAFunction(args)
	case "updateMSA":
		return nil, this.updateMSAFunction(args)
	case "submitMSA":
		return nil, this.submitMSAFunction(args)
	case "approveMSA":
		return nil, this.approveMSAFunction(args)
	//MSA Functions

	case "createSoWDO":
		return nil, this.createSoWDOFunction(args)
	case "updateDO":
		return nil, this.updateDOFunction(args)
	case "submitDO":
		return nil, this.submitDOFunction(args)
	case "approveDO":
		return nil, this.approveDOFunction(args)

	case "createSoWInvoice":
		return nil, this.createSoWInvoiceFunction(args)
	case "updateInvoice":
		return nil, this.updateInvoiceFunction(args)
	case "submitInvoice":
		return nil, this.submitInvoiceFunction(args)
	case "approveInvoice":
		return nil, this.approveInvoiceFunction(args)
	case "completeInvoice":
		return nil, this.completeInvoiceFunction(args)

	case "createExternalAttachmentToDoc":
		return nil, this.createExternalAttachmentToDocFunction(args)

	default:
		return nil, errors.New("UNKNOWN_INVOCATION|Received unknown function invocation")
	}
}

// Query callback representing the query of a chaincode
func (this *HDLS) QueryImpl(function string, args []string) (interface{}, error) {
	switch function {
	case "getInvokingStatus":
		return this.getInvokingStatus(args[0])
	case "queryStatus":
		return this.getInvokingStatus(args[0])
	case "jsontest":
		a := TestResponse{"foo", "bar"}
		return a, nil
	case "ping":
		return "pong2", nil
	// Cert
	case "metadata":
		return this.db.GetCallerMetadata()
	case "binding":
		return this.db.GetBinding()
	case "cert":
		tcertder, _ := this.db.GetCallerCertificate()
		tcert, _ := primitives.DERToX509Certificate(tcertder)
		return tcert, nil
	case "commonName":
		tcertder, _ := this.db.GetCallerCertificate()
		tcert, _ := primitives.DERToX509Certificate(tcertder)
		return tcert.Subject.CommonName, nil
	case "dump":
		return this.dump()
	case "errortest":
		return nil, errors.New("Error test")

	case "getSoW":
		return this.getSoWFunction(args)
	case "getSoWs":
		return this.getSoWsFunction(args)
	case "listSoWByTypeStatus":
		return this.listSoWByTypeStatusFunction(args)

	case "listApprovalStatussByBaseDocumentId":
		return this.listApprovalStatussByBaseDocumentIdFunction(args)
	case "testFunction":
		return this.testFunction()
	case "getUserRoleInfo":
		return this.getUserRoleInfo()
	case "listAttachmentMetadataByBaseDocumentId":
		return this.listAttachmentMetadataByBaseDocumentIdFunction(args)
	case "getBaseDocumentAttachment":
		return this.getBaseDocumentAttachmentFunction(args)

	//MSA Functions
	case "getMSA":
		return this.getMSAFunction(args)
	case "getMSAs":
		return this.getMSAsFunction(args)
	//MSA Functions

	//DO Functions
	case "getDO":
		return this.getDOFunction(args)
	case "listDOBySoW":
		return this.listDOBySoWFunction(args)
	case "listDOByInvoice":
		return this.listDOByInvoiceFunction(args)
	//DO Functions
	case "listInvoiceBySoW":
		return this.listInvoiceBySoWFunction(args)
	case "getInvoice":
		return this.getInvoiceFunction(args)

	case "getAttachmentMetadata":
		return this.getAttachmentMetadataFunction(args)

	case "getUserRoleInfoFromCert":
		return this.getUserRoleInfoFromCertFunction(args)

	default:
		return nil, errors.New("Received unknown function query")
	}

}
