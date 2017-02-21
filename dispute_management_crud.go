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


//List structures
type TransactionIdentifications struct {
	Data []TransactionIdentification	
}	
type TransactionInfos struct {
	Data []TransactionInfo	
}	
type Customers struct {
	Data []Customer	
}	
type Banks struct {
	Data []Bank	
}	
type Merchants struct {
	Data []Merchant	
}	
type PISPs struct {
	Data []PISP	
}	
type Resolutions struct {
	Data []Resolution	
}	
type ResolutionExecutions struct {
	Data []ResolutionExecution	
}	
type CustomerDisputes struct {
	Data []CustomerDispute	
}	

type Dump struct {
	Reference *References
	TransactionIdentification *TransactionIdentifications
	TransactionInfo *TransactionInfos
	Customer *Customers
	Bank *Banks
	Merchant *Merchants
	PISP *PISPs
	Resolution *Resolutions
	ResolutionExecution *ResolutionExecutions
	CustomerDispute *CustomerDisputes
}

func (this *HDLS) dump() (*Dump, error) {
	var d Dump
	var err error
	d.Reference, err = this.listReferences()
	if err != nil {
		return nil, err
	}
	d.TransactionIdentification, err = this.listTransactionIdentifications()
	if err != nil {
		return nil, err
	}
	d.TransactionInfo, err = this.listTransactionInfos()
	if err != nil {
		return nil, err
	}
	d.Customer, err = this.listCustomers()
	if err != nil {
		return nil, err
	}
	d.Bank, err = this.listBanks()
	if err != nil {
		return nil, err
	}
	d.Merchant, err = this.listMerchants()
	if err != nil {
		return nil, err
	}
	d.PISP, err = this.listPISPs()
	if err != nil {
		return nil, err
	}
	d.Resolution, err = this.listResolutions()
	if err != nil {
		return nil, err
	}
	d.ResolutionExecution, err = this.listResolutionExecutions()
	if err != nil {
		return nil, err
	}
	d.CustomerDispute, err = this.listCustomerDisputes()
	if err != nil {
		return nil, err
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
	if dump.TransactionIdentification != nil {
		for _, x := range dump.TransactionIdentification.Data {
			err = this.putTransactionIdentification(&x)
		}
	}
	if dump.TransactionInfo != nil {
		for _, x := range dump.TransactionInfo.Data {
			err = this.putTransactionInfo(&x)
		}
	}
	if dump.Customer != nil {
		for _, x := range dump.Customer.Data {
			err = this.putCustomer(&x)
		}
	}
	if dump.Bank != nil {
		for _, x := range dump.Bank.Data {
			err = this.putBank(&x)
		}
	}
	if dump.Merchant != nil {
		for _, x := range dump.Merchant.Data {
			err = this.putMerchant(&x)
		}
	}
	if dump.PISP != nil {
		for _, x := range dump.PISP.Data {
			err = this.putPISP(&x)
		}
	}
	if dump.Resolution != nil {
		for _, x := range dump.Resolution.Data {
			err = this.putResolution(&x)
		}
	}
	if dump.ResolutionExecution != nil {
		for _, x := range dump.ResolutionExecution.Data {
			err = this.putResolutionExecution(&x)
		}
	}
	if dump.CustomerDispute != nil {
		for _, x := range dump.CustomerDispute.Data {
			err = this.putCustomerDispute(&x)
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
		"Customer", 
		"Bank", 
		"Merchant", 
		"PISP", 
		"Resolution", 
		"ResolutionExecution", 
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


	//Remove all the referenced entities since they are already stored.
	
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


	//Remove all the referenced entities since they are already stored.
	
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

func (this *HDLS) listTransactionIdentifications() (*TransactionIdentifications, error) {
	this.logger.Infof("Call: listTransactionIdentification")

	rows, err := this.listAllRows("TransactionIdentification")
	if err != nil {
		return nil, err
	}

	var xs TransactionIdentifications
	for _, row := range rows {
		var x TransactionIdentification 
		if this.val(row, &x) == nil {
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
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
		return nil
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

func (this *HDLS) refIdTransactionInfoTransactionId(v string) string {
	return fmt.Sprintf("TransactionInfo.TransactionId=%v", v)
}

func (this *HDLS) putTransactionInfo(x *TransactionInfo) error {
	if x.Id == "" {
		x.Id, _ = this.idTransactionInfo(x)
	}

	dst := x	// copy


	//Remove all the referenced entities since they are already stored.
	
	err := this.putA("TransactionInfo", dst.Id, dst)
	if err != nil {
		return err
	}
	var ref *Reference
	var refId string
	refId = this.refIdTransactionInfoTransactionId(x.TransactionId)
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

func (this *HDLS) listTransactionInfos() (*TransactionInfos, error) {
	this.logger.Infof("Call: listTransactionInfo")

	rows, err := this.listAllRows("TransactionInfo")
	if err != nil {
		return nil, err
	}

	var xs TransactionInfos
	for _, row := range rows {
		var x TransactionInfo 
		if this.val(row, &x) == nil {
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}

func (this *HDLS) listTransactionInfosByTransactionId(v string) (*TransactionInfos, error) {

	var xs TransactionInfos
	refId := this.refIdTransactionInfoTransactionId(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getTransactionInfo(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
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
	var ref *Reference
	var refId string	

	curr, err := this.getTransactionInfo(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return nil
	}
	refId = this.refIdTransactionInfoTransactionId(curr.TransactionId)
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
// 4. CUSTOMER 
//------------------------

func (this *HDLS) refIdCustomerCustomerId(v string) string {
	return fmt.Sprintf("Customer.CustomerId=%v", v)
}
func (this *HDLS) refIdCustomerAccountId(v string) string {
	return fmt.Sprintf("Customer.AccountId=%v", v)
}

func (this *HDLS) putCustomer(x *Customer) error {
	if x.Id == "" {
		x.Id, _ = this.idCustomer(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
	
	err := this.putA("Customer", dst.Id, dst)
	if err != nil {
		return err
	}
	var ref *Reference
	var refId string
	refId = this.refIdCustomerCustomerId(x.CustomerId)
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
	
	refId = this.refIdCustomerAccountId(x.AccountId)
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

func (this *HDLS) getCustomer(id string) (*Customer, error) {
	this.logger.Infof("Call: getCustomer")

	var x Customer 
	err := this.getA("Customer", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listCustomers() (*Customers, error) {
	this.logger.Infof("Call: listCustomer")

	rows, err := this.listAllRows("Customer")
	if err != nil {
		return nil, err
	}

	var xs Customers
	for _, row := range rows {
		var x Customer 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}

func (this *HDLS) listCustomersByCustomerId(v string) (*Customers, error) {

	var xs Customers
	refId := this.refIdCustomerCustomerId(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getCustomer(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}
func (this *HDLS) listCustomersByAccountId(v string) (*Customers, error) {

	var xs Customers
	refId := this.refIdCustomerAccountId(v)
	reference, _ := this.getReference(refId)
	if reference != nil {
		for _, id := range reference.Ids {
			x, _ := this.getCustomer(id)
			if x != nil {
				xs.Data = append(xs.Data, *x)
			}
		}
	}

	return &xs, nil
}

func (this *HDLS) addCustomer(jsonStr string) error {

	var x Customer 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putCustomer(&x)
}

func (this *HDLS) idCustomer(x *Customer) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteCustomer(x *Customer) error {
	var err error
	var ref *Reference
	var refId string	

	curr, err := this.getCustomer(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return nil
	}
	refId = this.refIdCustomerCustomerId(curr.CustomerId)
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
	refId = this.refIdCustomerAccountId(curr.AccountId)
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
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
		}
	}
	err = this.delete("Customer", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteCustomer(x *Customer) error {
	if err := this.deleteCustomer(x); err != nil {
		return err
	}
	
	return this.putCustomer(x)
}
//------------------------
// 5. BANK 
//------------------------


func (this *HDLS) putBank(x *Bank) error {
	if x.Id == "" {
		x.Id, _ = this.idBank(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
	
	err := this.putA("Bank", dst.Id, dst)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) getBank(id string) (*Bank, error) {
	this.logger.Infof("Call: getBank")

	var x Bank 
	err := this.getA("Bank", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listBanks() (*Banks, error) {
	this.logger.Infof("Call: listBank")

	rows, err := this.listAllRows("Bank")
	if err != nil {
		return nil, err
	}

	var xs Banks
	for _, row := range rows {
		var x Bank 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}


func (this *HDLS) addBank(jsonStr string) error {

	var x Bank 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putBank(&x)
}

func (this *HDLS) idBank(x *Bank) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteBank(x *Bank) error {
	var err error
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
		}
	}
	err = this.delete("Bank", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteBank(x *Bank) error {
	if err := this.deleteBank(x); err != nil {
		return err
	}
	
	return this.putBank(x)
}
//------------------------
// 6. MERCHANT 
//------------------------


func (this *HDLS) putMerchant(x *Merchant) error {
	if x.Id == "" {
		x.Id, _ = this.idMerchant(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
	
	err := this.putA("Merchant", dst.Id, dst)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) getMerchant(id string) (*Merchant, error) {
	this.logger.Infof("Call: getMerchant")

	var x Merchant 
	err := this.getA("Merchant", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listMerchants() (*Merchants, error) {
	this.logger.Infof("Call: listMerchant")

	rows, err := this.listAllRows("Merchant")
	if err != nil {
		return nil, err
	}

	var xs Merchants
	for _, row := range rows {
		var x Merchant 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}


func (this *HDLS) addMerchant(jsonStr string) error {

	var x Merchant 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putMerchant(&x)
}

func (this *HDLS) idMerchant(x *Merchant) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteMerchant(x *Merchant) error {
	var err error
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
		}
	}
	err = this.delete("Merchant", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteMerchant(x *Merchant) error {
	if err := this.deleteMerchant(x); err != nil {
		return err
	}
	
	return this.putMerchant(x)
}
//------------------------
// 7. PISP 
//------------------------


func (this *HDLS) putPISP(x *PISP) error {
	if x.Id == "" {
		x.Id, _ = this.idPISP(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
	
	err := this.putA("PISP", dst.Id, dst)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) getPISP(id string) (*PISP, error) {
	this.logger.Infof("Call: getPISP")

	var x PISP 
	err := this.getA("PISP", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listPISPs() (*PISPs, error) {
	this.logger.Infof("Call: listPISP")

	rows, err := this.listAllRows("PISP")
	if err != nil {
		return nil, err
	}

	var xs PISPs
	for _, row := range rows {
		var x PISP 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}


func (this *HDLS) addPISP(jsonStr string) error {

	var x PISP 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putPISP(&x)
}

func (this *HDLS) idPISP(x *PISP) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deletePISP(x *PISP) error {
	var err error
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
		}
	}
	err = this.delete("PISP", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwritePISP(x *PISP) error {
	if err := this.deletePISP(x); err != nil {
		return err
	}
	
	return this.putPISP(x)
}
//------------------------
// 8. RESOLUTION 
//------------------------

func (this *HDLS) refIdResolutionOutcome(v string) string {
	return fmt.Sprintf("Resolution.Outcome=%v", v)
}

func (this *HDLS) putResolution(x *Resolution) error {
	if x.Id == "" {
		x.Id, _ = this.idResolution(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
	
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

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listResolutions() (*Resolutions, error) {
	this.logger.Infof("Call: listResolution")

	rows, err := this.listAllRows("Resolution")
	if err != nil {
		return nil, err
	}

	var xs Resolutions
	for _, row := range rows {
		var x Resolution 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
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
		return nil
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
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
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
// 9. RESOLUTIONEXECUTION 
//------------------------


func (this *HDLS) putResolutionExecution(x *ResolutionExecution) error {
	if x.Id == "" {
		x.Id, _ = this.idResolutionExecution(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
	
	err := this.putA("ResolutionExecution", dst.Id, dst)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) getResolutionExecution(id string) (*ResolutionExecution, error) {
	this.logger.Infof("Call: getResolutionExecution")

	var x ResolutionExecution 
	err := this.getA("ResolutionExecution", id, &x)
	if err != nil {
		this.logger.Infof("Error occured %v\n", err)
		return nil, err
	} else if x.Id == "" {
		return nil, nil
	}

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listResolutionExecutions() (*ResolutionExecutions, error) {
	this.logger.Infof("Call: listResolutionExecution")

	rows, err := this.listAllRows("ResolutionExecution")
	if err != nil {
		return nil, err
	}

	var xs ResolutionExecutions
	for _, row := range rows {
		var x ResolutionExecution 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
}


func (this *HDLS) addResolutionExecution(jsonStr string) error {

	var x ResolutionExecution 
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return err
	}

	return this.putResolutionExecution(&x)
}

func (this *HDLS) idResolutionExecution(x *ResolutionExecution) (string, error) {
	return x.Id, nil
}

func (this *HDLS) deleteResolutionExecution(x *ResolutionExecution) error {
	var err error
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
		}
	}
	err = this.delete("ResolutionExecution", x.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *HDLS) overwriteResolutionExecution(x *ResolutionExecution) error {
	if err := this.deleteResolutionExecution(x); err != nil {
		return err
	}
	
	return this.putResolutionExecution(x)
}
//------------------------
// 10. CUSTOMERDISPUTE 
//------------------------

func (this *HDLS) refIdCustomerDisputeStatus(v string) string {
	return fmt.Sprintf("CustomerDispute.Status=%v", v)
}

func (this *HDLS) putCustomerDispute(x *CustomerDispute) error {
	if x.Id == "" {
		x.Id, _ = this.idCustomerDispute(x)
	}

	dst := x	// copy

	//Save dst.TransactionInfo as a separate entity
	if dst.TransactionInfo != nil {
		if dst.TransactionInfo.Id == "" {
			dst.TransactionInfo.Id = dst.Id + "_TransactionInfo"			
		}
		err1 := this.putTransactionInfo(dst.TransactionInfo)
		if err1 != nil {
			return err1
		}
	}
	//Save dst.Customer as a separate entity
	if dst.Customer != nil {
		if dst.Customer.Id == "" {
			dst.Customer.Id = dst.Id + "_Customer"			
		}
		err1 := this.putCustomer(dst.Customer)
		if err1 != nil {
			return err1
		}
	}
	//Save dst.Bank as a separate entity
	if dst.Bank != nil {
		if dst.Bank.Id == "" {
			dst.Bank.Id = dst.Id + "_Bank"			
		}
		err1 := this.putBank(dst.Bank)
		if err1 != nil {
			return err1
		}
	}
	//Save dst.PISP as a separate entity
	if dst.PISP != nil {
		if dst.PISP.Id == "" {
			dst.PISP.Id = dst.Id + "_PISP"			
		}
		err1 := this.putPISP(dst.PISP)
		if err1 != nil {
			return err1
		}
	}
	//Save dst.Merchant as a separate entity
	if dst.Merchant != nil {
		if dst.Merchant.Id == "" {
			dst.Merchant.Id = dst.Id + "_Merchant"			
		}
		err1 := this.putMerchant(dst.Merchant)
		if err1 != nil {
			return err1
		}
	}
	//Save dst.Resolution as a separate entity
	if dst.Resolution != nil {
		if dst.Resolution.Id == "" {
			dst.Resolution.Id = dst.Id + "_Resolution"			
		}
		err1 := this.putResolution(dst.Resolution)
		if err1 != nil {
			return err1
		}
	}

	//Remove all the referenced entities since they are already stored.
	dst.TransactionInfo = nil
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

	x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
	if err != nil {
		return nil, err
	}
	x.Customer, err = this.getCustomer(x.Id + "_Customer")
	if err != nil {
		return nil, err
	}
	x.Bank, err = this.getBank(x.Id + "_Bank")
	if err != nil {
		return nil, err
	}
	x.PISP, err = this.getPISP(x.Id + "_PISP")
	if err != nil {
		return nil, err
	}
	x.Merchant, err = this.getMerchant(x.Id + "_Merchant")
	if err != nil {
		return nil, err
	}
	x.Resolution, err = this.getResolution(x.Id + "_Resolution")
	if err != nil {
		return nil, err
	}

	return &x, nil
}

func (this *HDLS) listCustomerDisputes() (*CustomerDisputes, error) {
	this.logger.Infof("Call: listCustomerDispute")

	rows, err := this.listAllRows("CustomerDispute")
	if err != nil {
		return nil, err
	}

	var xs CustomerDisputes
	for _, row := range rows {
		var x CustomerDispute 
		if this.val(row, &x) == nil {
			x.TransactionInfo, err = this.getTransactionInfo(x.Id + "_TransactionInfo")
			if err != nil {
				continue
			}
			x.Customer, err = this.getCustomer(x.Id + "_Customer")
			if err != nil {
				continue
			}
			x.Bank, err = this.getBank(x.Id + "_Bank")
			if err != nil {
				continue
			}
			x.PISP, err = this.getPISP(x.Id + "_PISP")
			if err != nil {
				continue
			}
			x.Merchant, err = this.getMerchant(x.Id + "_Merchant")
			if err != nil {
				continue
			}
			x.Resolution, err = this.getResolution(x.Id + "_Resolution")
			if err != nil {
				continue
			}
			xs.Data = append(xs.Data, x)
		}
	}
	return &xs, nil
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
	return x.Id, nil
}

func (this *HDLS) deleteCustomerDispute(x *CustomerDispute) error {
	var err error
	var ref *Reference
	var refId string	

	curr, err := this.getCustomerDispute(x.Id)
	if err != nil {
		return err
	} else if curr == nil {
		return nil
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
	//Delete x.TransactionInfo	
	if(x.TransactionInfo != nil) {
		err = this.deleteTransactionInfo(x.TransactionInfo)
		if err != nil {
			return err
		}
	}
	//Delete x.Customer	
	if(x.Customer != nil) {
		err = this.deleteCustomer(x.Customer)
		if err != nil {
			return err
		}
	}
	//Delete x.Bank	
	if(x.Bank != nil) {
		err = this.deleteBank(x.Bank)
		if err != nil {
			return err
		}
	}
	//Delete x.PISP	
	if(x.PISP != nil) {
		err = this.deletePISP(x.PISP)
		if err != nil {
			return err
		}
	}
	//Delete x.Merchant	
	if(x.Merchant != nil) {
		err = this.deleteMerchant(x.Merchant)
		if err != nil {
			return err
		}
	}
	//Delete x.Resolution	
	if(x.Resolution != nil) {
		err = this.deleteResolution(x.Resolution)
		if err != nil {
			return err
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
