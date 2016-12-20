package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

type MSAExt struct {
	Id string
	*MSA

	BaseDocument           BaseDocument
	AccessControlIndicator AccessControlIndicator
}

const IDPREF_MSA string = "MSA_"
const DOCTYPE_MSA string = "MSA"

func genMSAId(uuid string) string {
	return IDPREF_MSA + uuid
}

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
		docti, errti := this.putTransactionIdentification(disputeContent.TransactionIdentification)
		if errti != nil {
			return errti
		}
	} else {
		return errors.New("createDispute: Transaction Identification is missing")
	}

	if disputeContent.Customer != nil {
		disputeContent.Customer.Id = "Customer_" + uuid
		disputeContent.CustomerId = "Customer_" + uuid
		doccu, errcu := this.putCustomer(disputeContent.Customer)
		if errcu != nil {
			return errcu
		}
	}

	if disputeContent.Bank != nil {
		disputeContent.Bank.Id = "Bank_" + uuid
		disputeContent.BankId = "Bank_" + uuid
		docba, errba := this.putBank(disputeContent.Bank)
		if errba != nil {
			return errba
		}
	}

	if disputeContent.PISP != nil {
		disputeContent.PISP.Id = "PISP_" + uuid
		disputeContent.PISPId = "PISP_" + uuid
		docpi, errpi := this.putPISP(disputeContent.PISP)
		if errpi != nil {
			return errpi
		}
	}

	if disputeContent.Merchant != nil {
		disputeContent.Merchant.Id = "Merchant_" + uuid
		disputeContent.MerchantId = "Merchant_" + uuid
		docme, errme := this.putMerchant(disputeContent.Merchant)
		if errme != nil {
			return errme
		}
	}

	if disputeContent.Resolution != nil {
		disputeContent.Resolution.Id = "Resolution_" + uuid
		disputeContent.ResolutionId = "Resolution_" + uuid
		docre, errre := this.putResolution(disputeContent.Resolution)
		if errre != nil {
			return errre
		}
	}

	doccd, errcd := this.putCustomerDispute(disputeContent)
	if errcd != nil {
		return errcd
	}
	return nil
}

func (this *HDLS) getMSAFunction(args []string) (*MSAExt, error) {
	if len(args) != 1 {
		return nil, errors.New("getMSAFunction: number of argument is invalid.")
	}
	id := args[0]
	return this.getMSAExt(id)
}

func (this *HDLS) getMSAExt(id string) (*MSAExt, error) {

	this.logger.Debugf("getMSA")

	msa, errg := this.getMSA(id)
	if errg != nil {
		this.logger.Debugf("getMSA errg %s", errg)
		return nil, errg
	}
	if msa == nil {
		this.logger.Debugf("getMSA msa nil %s ", id)
		return nil, fmt.Errorf("ID not found: %s", id)
	}

	this.logger.Debugf("getMSA msa %s", msa)
	msaExt, errgr := this.getMSARelated(*msa)
	if errgr != nil {
		return nil, errgr
	}
	if !msaExt.AccessControlIndicator.Read {
		return nil, fmt.Errorf("Read not allowed for this user: %s", id)
	}
	return msaExt, errgr

}

func (this *HDLS) getMSARelated(msa MSA) (*MSAExt, error) {

	// BasicDocument
	//	basedoc, errgb := this.getBaseDocument(msa.BaseDocumentId)
	//	if errgb != nil {
	//		return nil, errgb
	//	}
	//	if basedoc == nil {
	//		return nil, fmt.Errorf("getMSARelated: basedoc not found %s", msa.BaseDocumentId)
	//	}
	//	// access control indicator
	//	as, errla := this.listApprovalStatussByBaseDocumentId(msa.BaseDocumentId)
	//	if errla != nil {
	//		return nil, errla
	//	}
	//	if as == nil || as.Data == nil {
	//		return nil, fmt.Errorf("getMSARelated: ApprovalStatus not found %s", msa.BaseDocumentId)
	//	}
	aobj, erraobj := this.getApprovalBaseDocument(msa.BaseDocumentId)
	if erraobj != nil {
		return nil, erraobj
	}
	//	aci, errca := this.calcAccessControlIndicator(basedoc, as.Data)
	aci, errca := this.calcAccessControlIndicator(aobj.BaseDocument, aobj.ApprovalStatuses)
	if errca != nil {
		return nil, errca
	}
	//	msaExt := MSAExt{msa.Id, &msa, *basedoc, *aci}
	msaExt := MSAExt{msa.Id, &msa, *(aobj.BaseDocument), *aci}
	return &msaExt, nil
}

func (this *HDLS) getMSAsFunction(args []string) ([]MSAExt, error) {
	if len(args) != 0 {
		return nil, errors.New("getMSAsFunction: number of argument is invalid.")
	}
	return this.getMSAs()
}

func (this *HDLS) getMSAs() ([]MSAExt, error) {

	// TODO: access control
	// TODO: ordering by date
	this.logger.Debugf("getMSAs")

	msas, err := this.listMSAs()
	if err != nil {
		this.logger.Debugf("getMSAs err %s", err)
		return nil, err
	}
	this.logger.Debugf("getMSAs msas %s", msas)
	var msaExts []MSAExt
	for _, msa := range msas.Data {
		msaExt, errgr := this.getMSARelated(msa)
		if errgr != nil {
			this.logger.Debugf("getMSAs errgr %s", errgr)
			return nil, errgr
		}
		if msaExt.AccessControlIndicator.Read {
			msaExts = append(msaExts, *msaExt)
		}
	}
	sort.Sort(sort.Reverse(MSAByCreationDate(msaExts)))
	return msaExts, nil
}
func (this *HDLS) updateMSAFunction(args []string) error {
	this.logger.Debugf("updateMSA")
	if len(args) != 1 {
		return errors.New("updateMSAFunction: number of argument is invalid.")
	}

	disputeContentJson := args[0]

	var disputeContent MSA = MSA{}
	err := json.Unmarshal([]byte(disputeContentJson), &disputeContent)
	if err != nil {
		return err
	}
	this.logger.Debugf("disputeContent: ", disputeContent)
	return this.updateMSA(disputeContent.Id, disputeContent)
}

// Convert the current object in DB to obj except for IDs and states.
// I.e., update customizable fields only. This method does not update DB.
func (this *HDLS) getModifiedMSA(id string, obj *MSA) (*MSA, error) {
	currObj, errCurr := this.getMSA(id)
	if errCurr != nil {
		return nil, errCurr
	}
	if currObj == nil {
		this.logger.Debugf("getModifiedMSA msa nil %s ", id)
		return nil, fmt.Errorf("ID not found: %s", id)
	}
	newObj := *obj
	newObj.Id = currObj.Id
	newObj.BaseDocumentId = currObj.BaseDocumentId
	newObj.SupplierId = currObj.SupplierId
	return &newObj, nil
}

func (this *HDLS) updateMSA(id string, disputeContent MSA) error {

	msa, errgs := this.getModifiedMSA(id, &disputeContent)
	if errgs != nil {
		return errgs
	}
	if msa == nil {
		this.logger.Debugf("updateMSA msa nil %s ", id)
		return fmt.Errorf("ID not found: %s", id)
	}

	// access control.
	errac := this.checkUpdatableOrError(msa.BaseDocumentId, msa.Id)
	if errac != nil {
		return errac
	}

	errpc := this.overwriteMSA(msa)
	if errpc != nil {
		return errpc
	}
	return nil
}

func (this *HDLS) submitMSAFunction(args []string) error {

	if len(args) != 1 {
		return errors.New("submitMSAFunction: number of argument is invalid.")
	}
	id := args[0]
	return this.submitMSA(id)
}

func (this *HDLS) submitMSA(id string) error {
	msa, errgs := this.getMSA(id)
	if errgs != nil {
		return errgs
	}
	if msa == nil {
		this.logger.Debugf("submitMSA msa nil %s ", id)
		return fmt.Errorf("ID not found: %s", id)
	}

	// access control.
	errac := this.checkUpdatableOrError(msa.BaseDocumentId, msa.Id)
	if errac != nil {
		return errac
	}

	//consistency checking
	var consistencyMsg = ""
	timeChecker := TimeChecker{BASEDOCUMENT_TIMELAYOUT}

	otherMsas, errmsas := this.getMSAs()
	if errmsas != nil {
		this.logger.Debugf("getMSAs error: %s", errmsas)
		return errmsas
	}
	if len(otherMsas) != 0 {
		for _, otherMsa := range otherMsas {

			if otherMsa.BaseDocument.Status == BASEDOCUMENTSTATUS_APPROVED {
				// compare supplierid
				if msa.SupplierId != otherMsa.MSA.SupplierId {
					consistencyMsg += fmt.Sprintf("SupplierId should be same. ")
					break
				}
				// compare buyerid
				if msa.ClientName != otherMsa.MSA.ClientName {
					consistencyMsg += fmt.Sprintf("ClientName should be same. ")
					break
				}
				// compare time
				//case 0 empty Date value
				if msa.EffectiveDate == "" || msa.ExpiryDate == "" {
					consistencyMsg += fmt.Sprintf("EffectiveDate or ExpiryDate's value is empty. ")
					break
				}
				//case 1 ExpiryDate should later than EffectiveDate
				msaTimeCmp, errMsaTimeCmp := timeChecker.CompareWithPeriod(msa.EffectiveDate, msa.ExpiryDate, msa.ExpiryDate)
				if errMsaTimeCmp != nil {
					return errMsaTimeCmp
				}
				if msaTimeCmp == +1 {
					consistencyMsg += fmt.Sprintf("ExpiryDate should later than EffectiveDate. ")
				}
				//case 2 startDate overlap other period
				msaTimeCmp, errMsaTimeCmp = timeChecker.CompareWithPeriod(msa.EffectiveDate, otherMsa.MSA.EffectiveDate, otherMsa.MSA.ExpiryDate)
				if errMsaTimeCmp != nil {
					return errMsaTimeCmp
				}
				if msaTimeCmp == 0 && msa.EffectiveDate != otherMsa.MSA.EffectiveDate && msa.EffectiveDate != otherMsa.MSA.ExpiryDate {
					consistencyMsg += fmt.Sprintf("EffectiveDate - ExpiryDate period overlap other MSA's period. ")
					break
				}
				//case 3 endDate overlap other period
				msaTimeCmp, errMsaTimeCmp = timeChecker.CompareWithPeriod(msa.ExpiryDate, otherMsa.MSA.EffectiveDate, otherMsa.MSA.ExpiryDate)
				if errMsaTimeCmp != nil {
					return errMsaTimeCmp
				}
				if msaTimeCmp == 0 && msa.ExpiryDate != otherMsa.MSA.EffectiveDate && msa.ExpiryDate != otherMsa.MSA.ExpiryDate {
					consistencyMsg += fmt.Sprintf("EffectiveDate - ExpiryDate period overlap other MSA's period. ")
					break
				}
				//case 4 period cotain other period
				msaTimeCmp, errMsaTimeCmp = timeChecker.CompareWithPeriod(otherMsa.MSA.EffectiveDate, msa.EffectiveDate, msa.ExpiryDate)
				if errMsaTimeCmp != nil {
					return errMsaTimeCmp
				}
				msaTimeCmp2, errMsaTimeCmp2 := timeChecker.CompareWithPeriod(otherMsa.MSA.ExpiryDate, msa.EffectiveDate, msa.ExpiryDate)
				if errMsaTimeCmp2 != nil {
					return errMsaTimeCmp2
				}
				if msaTimeCmp == 0 || msaTimeCmp2 == 0 {
					consistencyMsg += fmt.Sprintf("EffectiveDate - ExpiryDate period overlap other MSA's period. ")
					break
				}
			}
		}
	}
	if len(consistencyMsg) != 0 || len(msa.ConsistencyErrorMessage) != 0 {
		msa.ConsistencyErrorMessage = consistencyMsg
		errmsa := this.overwriteMSA(msa)
		if errmsa != nil {
			return errmsa
		}
	}
	if len(consistencyMsg) != 0 {
		return nil
	}
	errsubmit := this.submitBaseDocument(msa.BaseDocumentId, func(basedocid string) ([]ApprovalStatus, error) {
		approvalList := [1]ApprovalStatus{}
		approvalList[0].Init("", ROLE_BUYER, msa.ClientName, 0, basedocid)
		return approvalList[:], nil
	})

	if errsubmit != nil {
		return errsubmit
	}

	//publishing event
	errpublish := this.publishEvent(EVENTTYPE_SUBMITTED, DOCTYPE_MSA, id, msa.BaseDocumentId)
	if errpublish != nil {
		return errpublish
	}
	//publishing event

	return nil
}

func (this *HDLS) approveMSAFunction(args []string) error {

	this.logger.Debugf("approveMSAFunction %d", len(args))
	if len(args) != 3 {
		return errors.New("approveMSAFunction: number of argument is invalid.")
	}
	id := args[0]
	approval := args[1]
	comment := args[2]

	return this.approveMSA(id, approval, comment)
}

func (this *HDLS) approveMSA(id string, approval string, comment string) error {
	this.logger.Debugf("approveMSA %s, %s, %s", id, approval, comment)
	msa, errgs := this.getMSA(id)
	if errgs != nil {
		this.logger.Debugf("approveMSA errgs %s", errgs)
		return errgs
	}
	if msa == nil {
		this.logger.Debugf("approveMSA msa nil %s ", id)
		return fmt.Errorf("ID not found: %s", id)
	}

	// access control
	errac := this.checkApprovableOrError(msa.BaseDocumentId, msa.Id)
	if errac != nil {
		this.logger.Debugf("approveMSA errac %s", errac)
		return errac
	}
	isadmin, erria := this.isAdministrator()
	if erria != nil {
		return erria
	}

	errapprove := this.approveBaseDocument(msa.BaseDocumentId, approval, comment, isadmin)
	if errapprove != nil {
		return errapprove
	}

	//publishing event
	errpublish := this.publishEvent(approval, DOCTYPE_MSA, id, msa.BaseDocumentId)
	if errpublish != nil {
		return errpublish
	}
	//publishing event

	return nil
}

//////////////////////////////////////////////////////////////////////////
// sort

//func sortMSAByCreationDate(list []MSAExt) ([]MSAExt) {
//	var listSorted []*SortedDoc
//	for _, msa := range list {
//		listSorted = append(listSorted, &SortedDoc{ msa, msa.BaseDocument })
//	}
//	sort.Sort(sort.Reverse(ByCreationDate(listSorted)))
//	var listOut []MSAExt
//	for _, sdoc := range listSorted {
//		listOut = append(listOut, (*sdoc).DocRef.(MSAExt))
//	}
//	return listOut
//}

type MSAByCreationDate []MSAExt

func (s MSAByCreationDate) Len() int {
	return len(s)
}
func (s MSAByCreationDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s MSAByCreationDate) Less(i, j int) bool {
	iid := s[i].BaseDocument.CreationDate
	jid := s[j].BaseDocument.CreationDate
	return iid < jid
}
