package main

import (
	"encoding/json"
	"errors"
	"time"
)

func (hdls *HDLS) addNewCustomerDispute(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error

	if disputeContent.TransactionIdentification != nil {
		disputeContent.TransactionIdentification.Id = "TransactionIdentification_" + uuid
		disputeContent.TransactionIdentificationId = "TransactionIdentification_" + uuid
		err = hdls.putTransactionIdentification(disputeContent.TransactionIdentification)
		if err != nil {
			return err
		}
	} else {
		return errors.New("createDispute: Transaction Identification is missing")
	}

	if disputeContent.Customer != nil {
		disputeContent.Customer.Id = "Customer_" + uuid
		disputeContent.CustomerId = "Customer_" + uuid
		if disputeContent.Customer.TransactionInfo != nil {
			disputeContent.Customer.TransactionInfoId = "Customer_TxnInfo_" + uuid
			disputeContent.Customer.TransactionInfo.Id = "Customer_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.Customer.TransactionInfo)
			if err != nil {
				return err
			}
		}
		err = hdls.putCustomer(disputeContent.Customer)
		if err != nil {
			return err
		}
	}

	if disputeContent.Bank != nil {
		disputeContent.Bank.Id = "Bank_" + uuid
		disputeContent.BankId = "Bank_" + uuid
		if disputeContent.Bank.TransactionInfo != nil {
			disputeContent.Bank.TransactionInfoId = "Bank_TxnInfo_" + uuid
			disputeContent.Bank.TransactionInfo.Id = "Bank_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.Bank.TransactionInfo)
			if err != nil {
				return err
			}
		}
		err = hdls.putBank(disputeContent.Bank)
		if err != nil {
			return err
		}
	}

	if disputeContent.PISP != nil {
		disputeContent.PISP.Id = "PISP_" + uuid
		disputeContent.PISPId = "PISP_" + uuid
		if disputeContent.PISP.TransactionInfo != nil {
			disputeContent.PISP.TransactionInfoId = "PISP_TxnInfo_" + uuid
			disputeContent.PISP.TransactionInfo.Id = "PISP_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.PISP.TransactionInfo)
			if err != nil {
				return err
			}
		}
		err = hdls.putPISP(disputeContent.PISP)
		if err != nil {
			return err
		}
	}

	if disputeContent.Merchant != nil {
		disputeContent.Merchant.Id = "Merchant_" + uuid
		disputeContent.MerchantId = "Merchant_" + uuid
		if disputeContent.Merchant.TransactionInfo != nil {
			disputeContent.Merchant.TransactionInfoId = "Merchant_TxnInfo_" + uuid
			disputeContent.Merchant.TransactionInfo.Id = "Merchant_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.Merchant.TransactionInfo)
			if err != nil {
				return err
			}
		}
		err = hdls.putMerchant(disputeContent.Merchant)
		if err != nil {
			return err
		}
	}

	if disputeContent.Resolution != nil {
		disputeContent.Resolution.Id = "Resolution_" + uuid
		disputeContent.ResolutionId = "Resolution_" + uuid
		err = hdls.putResolution(disputeContent.Resolution)
		if err != nil {
			return err
		}
	}

	err = hdls.putCustomerDispute(&disputeContent)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) updateCustomerDispute(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	if disputeContent.TransactionIdentification != nil {
		if disputeContent.TransactionIdentification.Id == "" {
			disputeContent.TransactionIdentification.Id = "TransactionIdentification_" + uuid
			disputeContent.TransactionIdentificationId = "TransactionIdentification_" + uuid
			err = hdls.putTransactionIdentification(disputeContent.TransactionIdentification)
		} else {
			err = hdls.overwriteTransactionIdentification(disputeContent.TransactionIdentification)
		}

		if err != nil {
			return err
		}
	} else {
		return errors.New("createDispute: Transaction Identification is missing")
	}

	if disputeContent.Customer != nil {
		if disputeContent.Customer.Id == "" {
			disputeContent.Customer.Id = "Customer_" + uuid
			disputeContent.CustomerId = "Customer_" + uuid
			err = hdls.putCustomer(disputeContent.Customer)
		} else {
			err = hdls.overwriteCustomer(disputeContent.Customer)
		}
		if err != nil {
			return err
		}
	}

	if disputeContent.Bank != nil {
		if disputeContent.Bank.Id == "" {
			disputeContent.Bank.Id = "Bank_" + uuid
			disputeContent.BankId = "Bank_" + uuid
			err = hdls.putBank(disputeContent.Bank)
		} else {
			err = hdls.overwriteBank(disputeContent.Bank)
		}
		if err != nil {
			return err
		}
	}

	if disputeContent.PISP != nil {
		if disputeContent.PISP.Id == "" {
			disputeContent.PISP.Id = "PISP_" + uuid
			disputeContent.PISPId = "PISP_" + uuid
			err = hdls.putPISP(disputeContent.PISP)
		} else {
			err = hdls.overwritePISP(disputeContent.PISP)
		}
		if err != nil {
			return err
		}
	}

	if disputeContent.Merchant != nil {
		if disputeContent.Merchant.Id == "" {
			disputeContent.Merchant.Id = "Merchant_" + uuid
			disputeContent.MerchantId = "Merchant_" + uuid
			err = hdls.putMerchant(disputeContent.Merchant)
		} else {
			err = hdls.overwriteMerchant(disputeContent.Merchant)
		}
		if err != nil {
			return err
		}
	}

	if disputeContent.Resolution != nil {
		if disputeContent.Resolution.Id == "" {
			disputeContent.Resolution.Id = "Resolution_" + uuid
			disputeContent.ResolutionId = "Resolution_" + uuid
			err = hdls.putResolution(disputeContent.Resolution)
		} else {
			err = hdls.overwriteResolution(disputeContent.Resolution)
		}
		if err != nil {
			return err
		}
	}

	err = hdls.overwriteCustomerDispute(&disputeContent)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) updatePISPInformation(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("updatePISPAssignToMerchant: Existing dispute with id " + disputeContent.Id + " not found.")
	}

	found := false
	for _, element := range existingDispute.Owner {
		// element is the element from someSlice for where we are
		if element == "pisp" {
			found = true
		}
	}
	if found == false {
		return errors.New("You do not have ownership of this dispute")
	}

	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))

	if disputeContent.PISP != nil {
		disputeContent.PISP.Id = "PISP_" + uuid
		existingDispute.PISPId = "PISP_" + uuid
		if disputeContent.PISP.TransactionInfo != nil {
			disputeContent.PISP.TransactionInfoId = "PISP_TxnInfo_" + uuid
			disputeContent.PISP.TransactionInfo.Id = "PISP_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.PISP.TransactionInfo)
			if err != nil {
				return err
			}
		}
		err = hdls.putPISP(disputeContent.PISP)
		if err != nil {
			return err
		}
	}

	if disputeContent.Merchant != nil {
		if existingDispute.Merchant == nil {
			disputeContent.Merchant.Id = "Merchant_" + uuid
			existingDispute.MerchantId = "Merchant_" + uuid

			err = hdls.putMerchant(disputeContent.Merchant)
		} else {
			existingMerchant, err3 := hdls.getMerchant(existingDispute.Merchant.Id)
			if err3 != nil {
				return err3
			}
			existingMerchant.Name = disputeContent.Merchant.Name
			err = hdls.overwriteMerchant(existingMerchant)
		}
		if err != nil {
			return err
		}
	}

	existingDispute.Owner = []string{"merchant", "bank"}
	existingDispute.Status = disputeContent.Status
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) updateMerchantInformation(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("updateMerchantInformation: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	found := false
	var i int
	for index, element := range existingDispute.Owner {
		// element is the element from someSlice for where we are
		if element == "merchant" {
			found = true
			i = index
		}
	}
	if found == false {
		return errors.New("You do not have ownership of this dispute")
	}
	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))

	if disputeContent.Merchant != nil {
		if disputeContent.Merchant.TransactionInfo != nil {
			disputeContent.Merchant.TransactionInfo.Id = "Merchant_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.Merchant.TransactionInfo)
			if err != nil {
				return err
			}
		}
		if existingDispute.Merchant == nil {
			disputeContent.Merchant.Id = "Merchant_" + uuid
			existingDispute.MerchantId = "Merchant_" + uuid
			disputeContent.Merchant.TransactionInfoId = "Merchant_TxnInfo_" + uuid
			err = hdls.putMerchant(disputeContent.Merchant)
		} else {
			existingMerchant, err3 := hdls.getMerchant(existingDispute.Merchant.Id)
			if err3 != nil {
				return err3
			}
			existingMerchant.Name = disputeContent.Merchant.Name
			existingMerchant.Branch = disputeContent.Merchant.Branch
			existingMerchant.Terminal = disputeContent.Merchant.Terminal
			existingMerchant.Cashier = disputeContent.Merchant.Cashier
			existingMerchant.Receipts = disputeContent.Merchant.Receipts
			existingMerchant.TransactionInfoId = "Merchant_TxnInfo_" + uuid
			err = hdls.overwriteMerchant(existingMerchant)
		}
		if err != nil {
			return err
		}
	}

	existingDispute.Owner = removeElem(existingDispute.Owner, i)
	if len(existingDispute.Owner) == 0 {
		existingDispute.Status = "Waiting for Resolution"
		existingDispute.Owner = append(existingDispute.Owner, "pisp")
	} else {
		existingDispute.Status = disputeContent.Status
	}
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) sendToBankFromPISP(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("sendToBankFromPISP: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))

	if disputeContent.Bank != nil {
		if existingDispute.Bank == nil {
			disputeContent.Bank.Id = "Bank_" + uuid
			existingDispute.BankId = "Bank_" + uuid
			err = hdls.putBank(disputeContent.Bank)
		} else {
			existingBank, err3 := hdls.getBank(existingDispute.Bank.Id)
			if err3 != nil {
				return err3
			}
			existingBank.Name = disputeContent.Bank.Name
			err = hdls.overwriteBank(existingBank)
		}
		if err != nil {
			return err
		}
	}

	existingDispute.Owner = disputeContent.Owner
	existingDispute.Status = disputeContent.Status
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) updateBankInformation(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("updateBankInformation: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	found := false
	var i int
	for index, element := range existingDispute.Owner {
		// element is the element from someSlice for where we are
		if element == "bank" {
			found = true
			i = index
		}
	}
	if found == false {
		return errors.New("You do not have ownership of this dispute")
	}
	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))

	if disputeContent.Bank != nil {
		if disputeContent.Bank.TransactionInfo != nil {
			disputeContent.Bank.TransactionInfo.Id = "Bank_TxnInfo_" + uuid
			err = hdls.putTransactionInfo(disputeContent.Bank.TransactionInfo)
			if err != nil {
				return err
			}
		}
		if existingDispute.Bank == nil {
			disputeContent.Bank.Id = "Bank_" + uuid
			existingDispute.BankId = "Bank_" + uuid
			disputeContent.Bank.TransactionInfoId = "Bank_TxnInfo_" + uuid
			err = hdls.putBank(disputeContent.Bank)
		} else {
			existingBank, err3 := hdls.getBank(existingDispute.Bank.Id)
			if err3 != nil {
				return err3
			}
			existingBank.Name = disputeContent.Bank.Name
			existingBank.Branch = disputeContent.Bank.Branch
			existingBank.Terminal = disputeContent.Bank.Terminal
			existingBank.Cashier = disputeContent.Bank.Cashier
			existingBank.Receipts = disputeContent.Bank.Receipts
			existingBank.TransactionInfoId = "Bank_TxnInfo_" + uuid
			err = hdls.overwriteBank(existingBank)
		}
		if err != nil {
			return err
		}
	}

	existingDispute.Owner = removeElem(existingDispute.Owner, i)
	if len(existingDispute.Owner) == 0 {
		existingDispute.Status = "Waiting for Resolution"
		existingDispute.Owner = append(existingDispute.Owner, "pisp")
	} else {
		existingDispute.Status = disputeContent.Status
	}
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)

	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) proposeResolution(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("proposeResolution: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	found := false
	for _, element := range existingDispute.Owner {
		// element is the element from someSlice for where we are
		if element == "pisp" {
			found = true
		}
	}
	if found == false {
		return errors.New("You do not have ownership of this dispute")
	}
	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))

	if disputeContent.Resolution != nil {
		disputeContent.Resolution.Id = "Resolution_" + uuid
		disputeContent.Resolution.ResolutionTime = time.Now().Format(time.RFC850)
		existingDispute.ResolutionId = "Resolution_" + uuid
		disputeContent.Resolution.TransactionInfoId = "Resolution_Txn_" + uuid
		disputeContent.Resolution.TransactionInfo.Id = "Resolution_Txn_" + uuid
		err = hdls.putTransactionInfo(disputeContent.Resolution.TransactionInfo)
		if err != nil {
			return err
		}
		disputeContent.Resolution.TransactionInfo = nil
		err = hdls.putResolution(disputeContent.Resolution)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Resolution is missing")
	}

	existingDispute.Owner = []string{"bank", "merchant"}
	existingDispute.Status = "Resolution Proposed"
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) approveResolution(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("proposeResolution: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	found := false
	o := disputeContent.Owner[0]
	var index int
	for i, element := range existingDispute.Owner {
		// element is the element from someSlice for where we are
		if element == o {
			found = true
			index = i
		}
	}
	if found == false {
		return errors.New("You do not have ownership of this dispute")
	}
	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))
	existingDispute.Owner = removeElem(existingDispute.Owner, index)
	if len(existingDispute.Owner) == 0 {
		if existingDispute.Resolution.Type == "Against Customer" {
			existingDispute.Status = "Dispute Closed"
			existingDispute.Owner = []string{"customer"}
		} else {
			existingDispute.Status = "Executing Resolution"
			existingDispute.Owner = []string{"pisp", "merchant", "bank"}
		}
	}
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)
	if err != nil {
		return err
	}
	return nil
}

func (hdls *HDLS) rejectResolution(disputeContent CustomerDispute) error {
	stub := hdls.db
	uuid := stub.GetTxID()
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("proposeResolution: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	found := false
	o := disputeContent.Owner[0]
	for _, element := range existingDispute.Owner {
		// element is the element from someSlice for where we are
		if element == o {
			found = true
		}
	}
	if found == false {
		return errors.New("You do not have ownership of this dispute")
	}
	cp := existingDispute
	cp.Audit = nil
	b, e := json.Marshal(cp)
	if e != nil {
		return e
	}
	existingDispute.Audit = append(existingDispute.Audit, string(b))

	existingDispute.Owner = []string{"pisp"}
	existingDispute.Status = "Waiting for Resolution"
	existingDispute.LastUpdated = time.Now().Format(time.RFC850)

	err = hdls.overwriteCustomerDispute(existingDispute)
	if err != nil {
		return err
	}
	return nil
}

func removeElem(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
