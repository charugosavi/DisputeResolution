/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

/* Do not modify this source code directly.
 * This is automatically generated by model_crud.go.php */

package main

import (
	"encoding/json"
	"fmt"
	"errors"
)

var _ = fmt.Printf
var _ = errors.New("Temp")

type Dump struct {
	Reference *References
}

func (this *HDLS) dump() (*Dump, error) {
	var d Dump
	var err error
	{
		d.Reference, err = this.listReferences()
		if err != nil {
				return nil, err
		}
	}

	return &d, err
}

func (this *HDLS) imprt(dump *Dump) (error) {
	var err error
	if dump.Reference != nil {
		for _, x := range dump.Reference.Data {
			err = this.putReference(&x)
		}
	}
	return err
}

func (this *HDLS) imprtJson(jsonStr string) (error) {
	var d Dump
	err := json.Unmarshal([]byte(jsonStr), &d)
	if err != nil {
		return err
	}

	return this.imprt(&d)
}

func (this *HDLS) createSchema() {
	models := []string{
		"Reference", 
		"TransactionIdentification", 
		"TransactionInfo", 
		"InvolvedParty", 
		"Resolution", 
		"CustomerDispute", 
	}
	for _, model := range models {
		this.createKeyEntTable(model)
	}
}
//------------------------
// 1. REFERENCE 
//------------------------


func (this *HDLS) putReference(x *Reference) error {
	if x.Id == "" {
		x.Id, _ = this.idReference(x)
	}

	dst := x	// copy
	
	err := this.putA("Reference", dst.Id, dst)
	if err != nil {
		return err
	}


	return nil
}

func (this *HDLS) getReference(id string) (*Reference, error) {
	this.logger.Infof("Call: getReference")

	var x Reference 
	err := this.getA("Reference", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}


	return &x, nil
}

func (this *HDLS) listReferences() (*References, error) {
	this.logger.Infof("Call: listReference")

	rows, err := this.listAllRows("Reference")
	if err != nil {
		return nil, err
	}

	var xs References
	for _, row := range rows {
		var x Reference 
		if this.val(row, &x) == nil {
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}


func (this *HDLS) addReference(jsonStr string) error {

	var x Reference 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putReference(&x)
}

func (this *HDLS) idReference(x *Reference) (string, error) {
	return this.db.GetTxID(), nil
}

func (this *HDLS) deleteReference(x *Reference) error {

	var err error



	err = this.delete("Reference", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteReference(x *Reference) error {
	if err := this.deleteReference(x); err != nil {
		return err
	}
	
	return this.putReference(x)
}
//------------------------
// 2. TRANSACTIONIDENTIFICATION 
//------------------------

func (this *HDLS) refIdTransactionIdentificationCustomerId(v string) string {
	return fmt.Sprintf("TransactionIdentification.CustomerId=%v", v)
}
func (this *HDLS) refIdTransactionIdentificationAccountId(v string) string {
	return fmt.Sprintf("TransactionIdentification.AccountId=%v", v)
}

func (this *HDLS) putTransactionIdentification(x *TransactionIdentification) error {
	if x.Id == "" {
		x.Id, _ = this.idTransactionIdentification(x)
	}

	dst := x	// copy
	
	err := this.putA("TransactionIdentification", dst.Id, dst)
	if err != nil {
		return err
	}

	var ref *Reference
	var refId string
	refId = this.refIdTransactionIdentificationCustomerId(x.CustomerId)
	ref, _ = this.getReference(refId)
	if ref == nil {
		ref = &Reference{
			Id : refId,
			Ids: []string{dst.Id},
		}
		err = this.putReference(ref)
	} else {
		ref.Ids = append(ref.Ids, dst.Id)
		err = this.overwriteReference(ref)
	}
	if err != nil {
		return err
	}
	
	refId = this.refIdTransactionIdentificationAccountId(x.AccountId)
	ref, _ = this.getReference(refId)
	if ref == nil {
		ref = &Reference{
			Id : refId,
			Ids: []string{dst.Id},
		}
		err = this.putReference(ref)
	} else {
		ref.Ids = append(ref.Ids, dst.Id)
		err = this.overwriteReference(ref)
	}
	if err != nil {
		return err
	}
	

	return nil
}

func (this *HDLS) getTransactionIdentification(id string) (*TransactionIdentification, error) {
	this.logger.Infof("Call: getTransactionIdentification")

	var x TransactionIdentification 
	err := this.getA("TransactionIdentification", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}


	return &x, nil
}


func (this *HDLS) listTransactionIdentificationsByCustomerId(v string) (*TransactionIdentifications, error) {

	var xs TransactionIdentifications
	refId := this.refIdTransactionIdentificationCustomerId(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getTransactionIdentification(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}
func (this *HDLS) listTransactionIdentificationsByAccountId(v string) (*TransactionIdentifications, error) {

	var xs TransactionIdentifications
	refId := this.refIdTransactionIdentificationAccountId(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getTransactionIdentification(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}

func (this *HDLS) addTransactionIdentification(jsonStr string) error {

	var x TransactionIdentification 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putTransactionIdentification(&x)
}

func (this *HDLS) idTransactionIdentification(x *TransactionIdentification) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteTransactionIdentification(x *TransactionIdentification) error {

	var err error

	var ref *Reference
	var refId string

	curr, err := this.getTransactionIdentification(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return errors.New("NOT_FOUND")
	}

	refId = this.refIdTransactionIdentificationCustomerId(curr.CustomerId)
	ref, _ = this.getReference(refId)
	if ref != nil {
		ref.Ids = remove(ref.Ids, x.Id)
		this.deleteReference(ref)
		
		if len(ref.Ids) > 0 {
			this.overwriteReference(ref)
		} else {
			this.deleteReference(ref)
		}
	}
	refId = this.refIdTransactionIdentificationAccountId(curr.AccountId)
	ref, _ = this.getReference(refId)
	if ref != nil {
		ref.Ids = remove(ref.Ids, x.Id)
		this.deleteReference(ref)
		
		if len(ref.Ids) > 0 {
			this.overwriteReference(ref)
		} else {
			this.deleteReference(ref)
		}
	}

	err = this.delete("TransactionIdentification", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteTransactionIdentification(x *TransactionIdentification) error {
	if err := this.deleteTransactionIdentification(x); err != nil {
		return err
	}
	
	return this.putTransactionIdentification(x)
}
//------------------------
// 3. TRANSACTIONINFO 
//------------------------


func (this *HDLS) putTransactionInfo(x *TransactionInfo) error {
	if x.Id == "" {
		x.Id, _ = this.idTransactionInfo(x)
	}

	dst := x	// copy
	
	err := this.putA("TransactionInfo", dst.Id, dst)
	if err != nil {
		return err
	}


	return nil
}

func (this *HDLS) getTransactionInfo(id string) (*TransactionInfo, error) {
	this.logger.Infof("Call: getTransactionInfo")

	var x TransactionInfo 
	err := this.getA("TransactionInfo", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}


	return &x, nil
}



func (this *HDLS) addTransactionInfo(jsonStr string) error {

	var x TransactionInfo 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putTransactionInfo(&x)
}

func (this *HDLS) idTransactionInfo(x *TransactionInfo) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteTransactionInfo(x *TransactionInfo) error {

	var err error



	err = this.delete("TransactionInfo", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteTransactionInfo(x *TransactionInfo) error {
	if err := this.deleteTransactionInfo(x); err != nil {
		return err
	}
	
	return this.putTransactionInfo(x)
}
//------------------------
// 4. INVOLVEDPARTY 
//------------------------


func (this *HDLS) putInvolvedParty(x *InvolvedParty) error {
	if x.Id == "" {
		x.Id, _ = this.idInvolvedParty(x)
	}

	dst := x	// copy
	dst.Transaction = nil
	
	err := this.putA("InvolvedParty", dst.Id, dst)
	if err != nil {
		return err
	}


	return nil
}

func (this *HDLS) getInvolvedParty(id string) (*InvolvedParty, error) {
	this.logger.Infof("Call: getInvolvedParty")

	var x InvolvedParty 
	err := this.getA("InvolvedParty", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.Transaction, err = this.getTransaction(x.TransactionId)
	if err != nil {
		return nil, err
	}

	return &x, nil
}



func (this *HDLS) addInvolvedParty(jsonStr string) error {

	var x InvolvedParty 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putInvolvedParty(&x)
}

func (this *HDLS) idInvolvedParty(x *InvolvedParty) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteInvolvedParty(x *InvolvedParty) error {

	var err error



	err = this.delete("InvolvedParty", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteInvolvedParty(x *InvolvedParty) error {
	if err := this.deleteInvolvedParty(x); err != nil {
		return err
	}
	
	return this.putInvolvedParty(x)
}
//------------------------
// 5. RESOLUTION 
//------------------------

func (this *HDLS) refIdResolutionOutcome(v string) string {
	return fmt.Sprintf("Resolution.Outcome=%v", v)
}

func (this *HDLS) putResolution(x *Resolution) error {
	if x.Id == "" {
		x.Id, _ = this.idResolution(x)
	}

	dst := x	// copy
	dst.Transaction = nil
	
	err := this.putA("Resolution", dst.Id, dst)
	if err != nil {
		return err
	}

	var ref *Reference
	var refId string
	refId = this.refIdResolutionOutcome(x.Outcome)
	ref, _ = this.getReference(refId)
	if ref == nil {
		ref = &Reference{
			Id : refId,
			Ids: []string{dst.Id},
		}
		err = this.putReference(ref)
	} else {
		ref.Ids = append(ref.Ids, dst.Id)
		err = this.overwriteReference(ref)
	}
	if err != nil {
		return err
	}
	

	return nil
}

func (this *HDLS) getResolution(id string) (*Resolution, error) {
	this.logger.Infof("Call: getResolution")

	var x Resolution 
	err := this.getA("Resolution", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.Transaction, err = this.getTransaction(x.TransactionId)
	if err != nil {
		return nil, err
	}

	return &x, nil
}


func (this *HDLS) listResolutionsByOutcome(v string) (*Resolutions, error) {

	var xs Resolutions
	refId := this.refIdResolutionOutcome(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getResolution(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}

func (this *HDLS) addResolution(jsonStr string) error {

	var x Resolution 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putResolution(&x)
}

func (this *HDLS) idResolution(x *Resolution) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteResolution(x *Resolution) error {

	var err error

	var ref *Reference
	var refId string

	curr, err := this.getResolution(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return errors.New("NOT_FOUND")
	}

	refId = this.refIdResolutionOutcome(curr.Outcome)
	ref, _ = this.getReference(refId)
	if ref != nil {
		ref.Ids = remove(ref.Ids, x.Id)
		this.deleteReference(ref)
		
		if len(ref.Ids) > 0 {
			this.overwriteReference(ref)
		} else {
			this.deleteReference(ref)
		}
	}

	err = this.delete("Resolution", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteResolution(x *Resolution) error {
	if err := this.deleteResolution(x); err != nil {
		return err
	}
	
	return this.putResolution(x)
}
//------------------------
// 6. CUSTOMERDISPUTE 
//------------------------

func (this *HDLS) refIdCustomerDisputeStatus(v string) string {
	return fmt.Sprintf("CustomerDispute.Status=%v", v)
}

func (this *HDLS) putCustomerDispute(x *CustomerDispute) error {
	if x.Id == "" {
		x.Id, _ = this.idCustomerDispute(x)
	}

	dst := x	// copy
	dst.Transaction = nil
	dst.Customer = nil
	dst.Bank = nil
	dst.PISP = nil
	dst.Merchant = nil
	dst.Resolution = nil
	
	err := this.putA("CustomerDispute", dst.Id, dst)
	if err != nil {
		return err
	}

	var ref *Reference
	var refId string
	refId = this.refIdCustomerDisputeStatus(x.Status)
	ref, _ = this.getReference(refId)
	if ref == nil {
		ref = &Reference{
			Id : refId,
			Ids: []string{dst.Id},
		}
		err = this.putReference(ref)
	} else {
		ref.Ids = append(ref.Ids, dst.Id)
		err = this.overwriteReference(ref)
	}
	if err != nil {
		return err
	}
	

	return nil
}

func (this *HDLS) getCustomerDispute(id string) (*CustomerDispute, error) {
	this.logger.Infof("Call: getCustomerDispute")

	var x CustomerDispute 
	err := this.getA("CustomerDispute", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.Transaction, err = this.getTransaction(x.TransactionId)
	if err != nil {
		return nil, err
	}
	x.Customer, err = this.getCustomer(x.CustomerId)
	if err != nil {
		return nil, err
	}
	x.Bank, err = this.getBank(x.BankId)
	if err != nil {
		return nil, err
	}
	x.PISP, err = this.getPISP(x.PISPId)
	if err != nil {
		return nil, err
	}
	x.Merchant, err = this.getMerchant(x.MerchantId)
	if err != nil {
		return nil, err
	}
	x.Resolution, err = this.getResolution(x.ResolutionId)
	if err != nil {
		return nil, err
	}

	return &x, nil
}


func (this *HDLS) listCustomerDisputesByStatus(v string) (*CustomerDisputes, error) {

	var xs CustomerDisputes
	refId := this.refIdCustomerDisputeStatus(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getCustomerDispute(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}

func (this *HDLS) addCustomerDispute(jsonStr string) error {

	var x CustomerDispute 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putCustomerDispute(&x)
}

func (this *HDLS) idCustomerDispute(x *CustomerDispute) (string, error) {
	return x.DisputeId, nil
}

func (this *HDLS) deleteCustomerDispute(x *CustomerDispute) error {

	var err error

	var ref *Reference
	var refId string

	curr, err := this.getCustomerDispute(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return errors.New("NOT_FOUND")
	}

	refId = this.refIdCustomerDisputeStatus(curr.Status)
	ref, _ = this.getReference(refId)
	if ref != nil {
		ref.Ids = remove(ref.Ids, x.Id)
		this.deleteReference(ref)
		
		if len(ref.Ids) > 0 {
			this.overwriteReference(ref)
		} else {
			this.deleteReference(ref)
		}
	}

	err = this.delete("CustomerDispute", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteCustomerDispute(x *CustomerDispute) error {
	if err := this.deleteCustomerDispute(x); err != nil {
		return err
	}
	
	return this.putCustomerDispute(x)
}