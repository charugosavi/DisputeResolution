package main

import (
	"encoding/json"
	"errors"
	"time"
)

func (hdls *HDLS) addNewCustomerDispute(disputeContent CustomerDispute) error {
	return hdls.putCustomerDispute(&disputeContent)
}

func (hdls *HDLS) updatePISPInformation(disputeContent CustomerDispute) error {
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
	existingDispute.PISP = disputeContent.PISP
	existingDispute.Merchant = disputeContent.Merchant
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
	existingDispute.Merchant = disputeContent.Merchant
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

func (hdls *HDLS) updateBankInformation(disputeContent CustomerDispute) error {
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
	existingDispute.Bank = disputeContent.Bank
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
	existingDispute.Resolution = disputeContent.Resolution
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
		if existingDispute.Resolution.Outcome == "Against Customer" {
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

func (hdls *HDLS) executeResolution(disputeContent CustomerDispute) error {
	var err error
	existingDispute, err2 := hdls.getCustomerDispute(disputeContent.Id)

	if err2 != nil {
		return err2
	}

	if existingDispute == nil {
		return errors.New("proposeResolution: Existing dispute with id " + disputeContent.Id + " not found.")
	}
	found := false
	var index int
	o := disputeContent.Owner[0]
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
	existingDispute.ResolutionExecution = append(existingDispute.ResolutionExecution, disputeContent.ResolutionExecution[0])

	existingDispute.Owner = removeElem(existingDispute.Owner, index)
	if len(existingDispute.Owner) == 0 {
		existingDispute.Status = "Dispute Closed"
		existingDispute.Owner = []string{"customer"}
	}
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
