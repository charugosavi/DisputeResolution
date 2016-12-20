package main

import (
	"encoding/json"
	"errors"
)

func (this *HDLS) addCustomerDisputeFunction(args []string) error {
	this.logger.Debugf("createDispute")
	if len(args) != 1 {
		return errors.New("createDispute: number of argument is invalid.")
	}
	disputeContentJson := args[0]
	var disputeContent CustomerDispute = CustomerDispute{}
	err := json.Unmarshal([]byte(disputeContentJson), &disputeContent)
	if err != nil {
		return err
	}
	this.logger.Debugf("disputeContent: ", disputeContent)
	return this.createDispute(disputeContent)
}

func (this *HDLS) createDispute(disputeContent CustomerDispute) error {
	stub := this.db
	uuid := stub.GetTxID()

	if disputeContent.TransactionIdentification != nil {
		disputeContent.TransactionIdentification.Id = "TransactionIdentification_" + uuid
		disputeContent.TransactionIdentificationId = "TransactionIdentification_" + uuid
		errti := this.putTransactionIdentification(disputeContent.TransactionIdentification)
		if errti != nil {
			return errti
		}
	} else {
		return errors.New("createDispute: Transaction Identification is missing")
	}

	if disputeContent.Customer != nil {
		disputeContent.Customer.Id = "Customer_" + uuid
		disputeContent.CustomerId = "Customer_" + uuid
		errcu := this.putCustomer(disputeContent.Customer)
		if errcu != nil {
			return errcu
		}
	}

	if disputeContent.Bank != nil {
		disputeContent.Bank.Id = "Bank_" + uuid
		disputeContent.BankId = "Bank_" + uuid
		errba := this.putBank(disputeContent.Bank)
		if errba != nil {
			return errba
		}
	}

	if disputeContent.PISP != nil {
		disputeContent.PISP.Id = "PISP_" + uuid
		disputeContent.PISPId = "PISP_" + uuid
		errpi := this.putPISP(disputeContent.PISP)
		if errpi != nil {
			return errpi
		}
	}

	if disputeContent.Merchant != nil {
		disputeContent.Merchant.Id = "Merchant_" + uuid
		disputeContent.MerchantId = "Merchant_" + uuid
		errme := this.putMerchant(disputeContent.Merchant)
		if errme != nil {
			return errme
		}
	}

	if disputeContent.Resolution != nil {
		disputeContent.Resolution.Id = "Resolution_" + uuid
		disputeContent.ResolutionId = "Resolution_" + uuid
		errre := this.putResolution(disputeContent.Resolution)
		if errre != nil {
			return errre
		}
	}

	errcd := this.putCustomerDispute(&disputeContent)
	if errcd != nil {
		return errcd
	}
	return nil
}
