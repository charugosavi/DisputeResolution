package main

import (
	"encoding/json"
	"errors"
)

func (hdls *HDLS) addCustomerDisputeFunction(args []string) error {
	hdls.logger.Debugf("addCustomerDisputeFunction")
	if len(args) != 1 {
		return errors.New("addCustomerDisputeFunction: number of argument is invalid")
	}
	disputeContentJSON := args[0]
	disputeContent := CustomerDispute{}
	err := json.Unmarshal([]byte(disputeContentJSON), &disputeContent)
	if err != nil {
		return err
	}
	hdls.logger.Debugf("disputeContent: ", disputeContent)
	return hdls.createDispute(disputeContent)
}

func (hdls *HDLS) updateCustomerDisputeFunction(args []string) error {
	hdls.logger.Debugf("updateCustomerDisputeFunction")
	if len(args) != 1 {
		return errors.New("updateCustomerDisputeFunction: number of argument is invalid")
	}
	disputeContentJSON := args[0]
	disputeContent := CustomerDispute{}
	err := json.Unmarshal([]byte(disputeContentJSON), &disputeContent)
	if err != nil {
		return err
	}
	hdls.logger.Debugf("disputeContent: ", disputeContent)
	return hdls.updateDispute(disputeContent)
}

func (hdls *HDLS) createDispute(disputeContent CustomerDispute) error {
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
		err = hdls.putCustomer(disputeContent.Customer)
		if err != nil {
			return err
		}
	}

	if disputeContent.Bank != nil {
		disputeContent.Bank.Id = "Bank_" + uuid
		disputeContent.BankId = "Bank_" + uuid
		err = hdls.putBank(disputeContent.Bank)
		if err != nil {
			return err
		}
	}

	if disputeContent.PISP != nil {
		disputeContent.PISP.Id = "PISP_" + uuid
		disputeContent.PISPId = "PISP_" + uuid
		err = hdls.putPISP(disputeContent.PISP)
		if err != nil {
			return err
		}
	}

	if disputeContent.Merchant != nil {
		disputeContent.Merchant.Id = "Merchant_" + uuid
		disputeContent.MerchantId = "Merchant_" + uuid
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

func (hdls *HDLS) updateDispute(disputeContent CustomerDispute) error {
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
