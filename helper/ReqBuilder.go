package helper

import "BNI-TAP/model"

func ParseInqDB() {
	var ReqPayBNI model.ReqPayBNI
	var ResInqBNI model.ResInqBni
	ReqPayBNI.AdditionalLabel1 = ResInqBNI.AdditionalLabel1
	ReqPayBNI.AdditionalLabel2 = ResInqBNI.AdditionalLabel2
	ReqPayBNI.AdditionalLabel3 = ResInqBNI.AdditionalLabel3
	ReqPayBNI.AdditionalValue1 = ResInqBNI.AdditionalValue1
	ReqPayBNI.AdditionalValue2 = ResInqBNI.AdditionalValue2
	ReqPayBNI.AdditionalValue3 = ResInqBNI.AdditionalValue3
	ReqPayBNI.Amount = ResInqBNI.BilledAmount + ResInqBNI.FeeAmount
	ReqPayBNI.BilledAmount = ResInqBNI.BilledAmount
	ReqPayBNI.BilledAmountLabel = ResInqBNI.BilledAmountLabel
	ReqPayBNI.BilledAmountValue = ResInqBNI.BilledAmountValue
	ReqPayBNI.BillingNumber = ResInqBNI.BillingNumber
	ReqPayBNI.Channel = "TELLER"
	ReqPayBNI.ClientID = ResInqBNI.ClientID
	ReqPayBNI.Currency = ResInqBNI.Currency
	ReqPayBNI.FeeAmount = ResInqBNI.FeeAmount
	ReqPayBNI.FeeAmountLabel = ResInqBNI.FeeAmountLabel
	ReqPayBNI.FeeAmountValue = ResInqBNI.FeeAmountValue
	ReqPayBNI.LanguageID = "ID"
	ReqPayBNI.PaymentType = ResInqBNI.BilledAmountValue
	ReqPayBNI.TrxID = ResInqBNI.TrxID
	ReqPayBNI.VaNameLabel = ResInqBNI.VaNameLabel
	ReqPayBNI.VirtualAccountName = ResInqBNI.VirtualAccountName
	ReqPayBNI.VirtualAccountNumber = ResInqBNI.VirtualAccountNumber
	ReqPayBNI.VirtualAccountTrxType = ResInqBNI.VirtualAccountTrxType
}
func SignReqPay(reqpay model.ReqPayBNI, jwt string) model.SignedReqPayBNI {
	var SignedReqPayBNI model.SignedReqPayBNI
	var ReqPayBNI model.ReqPayBNI
	SignedReqPayBNI.AdditionalLabel1 = ReqPayBNI.AdditionalLabel1
	SignedReqPayBNI.AdditionalLabel2 = ReqPayBNI.AdditionalLabel2
	SignedReqPayBNI.AdditionalLabel3 = ReqPayBNI.AdditionalLabel3
	SignedReqPayBNI.AdditionalValue1 = ReqPayBNI.AdditionalValue1
	SignedReqPayBNI.AdditionalValue2 = ReqPayBNI.AdditionalValue2
	SignedReqPayBNI.AdditionalValue3 = ReqPayBNI.AdditionalValue3
	SignedReqPayBNI.Amount = ReqPayBNI.Amount
	SignedReqPayBNI.BilledAmount = ReqPayBNI.BilledAmount
	SignedReqPayBNI.BilledAmountLabel = ReqPayBNI.BilledAmountLabel
	SignedReqPayBNI.BilledAmountValue = ReqPayBNI.BilledAmountValue
	SignedReqPayBNI.BillingNumber = ReqPayBNI.BillingNumber
	SignedReqPayBNI.Channel = ReqPayBNI.Channel
	SignedReqPayBNI.ClientID = ReqPayBNI.ClientID
	SignedReqPayBNI.Currency = ReqPayBNI.Currency
	SignedReqPayBNI.FeeAmount = ReqPayBNI.FeeAmount
	SignedReqPayBNI.FeeAmountLabel = ReqPayBNI.FeeAmountLabel
	SignedReqPayBNI.FeeAmountValue = ReqPayBNI.FeeAmountValue
	SignedReqPayBNI.LanguageID = ReqPayBNI.LanguageID
	SignedReqPayBNI.PaymentType = ReqPayBNI.PaymentType
	SignedReqPayBNI.TrxID = ReqPayBNI.TrxID
	SignedReqPayBNI.VaNameLabel = ReqPayBNI.VaNameLabel
	SignedReqPayBNI.VirtualAccountName = ReqPayBNI.VirtualAccountName
	SignedReqPayBNI.VirtualAccountNumber = ReqPayBNI.VirtualAccountNumber
	SignedReqPayBNI.VirtualAccountTrxType = ReqPayBNI.VirtualAccountTrxType
	SignedReqPayBNI.Signature = jwt

	return SignedReqPayBNI
}
