package model

type ResPayBNI struct {
	Amount                string `json:"Amount"`
	BillingNumber         string `json:"billingNumber"`
	BillingLabel          string `json:"billingLabel"`
	ClientID              string `json:"clientId"`
	TrxID                 string `json:"trxId"`
	VirtualAccountNumber  string `json:"virtualAccountNumber"`
	VirtualAccountName    string `json:"virtualAccountName"`
	VaNameLabel           string `json:"vaNameLabel"`
	VirtualAccountTrxType string `json:"virtualAccountTrxType"`
	BilledAmount          string `json:"billedAmount"`
	BilledAmountLabel     string `json:"billedAmountLabel"`
	BilledAmountValue     string `json:"billedAmountValue"`
	AdditionalLabel1      string `json:"additionalLabel1"`
	AdditionalLabel2      string `json:"additionalLabel2"`
	AdditionalLabel3      string `json:"additionalLabel3"`
	AdditionalValue1      string `json:"additionalValue1"`
	AdditionalValue2      string `json:"additionalValue2"`
	AdditionalValue3      string `json:"additionalValue3"`
	FeeAmount             string `json:"feeAmount"`
	FeeAmountLabel        string `json:"feeAmountLabel"`
	FeeAmountValue        string `json:"feeAmountValue"`
	Currency              string `json:"currency"`
	Channel               string `json:"channel"`
	PaymentType           string `json:"paymentType"`
	JournalNum            string `json:"journalNum"`
	TrxDateTime           string `json:"trxDateTime"`
}
