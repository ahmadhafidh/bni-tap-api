package model

type ReqInqBNI struct {
	BillingNumber string `json:"billingNumber"`
	LanguageId    string `json:"languageId"`
}

type SignedReqInqBNI struct {
	BillingNumber string `json:"billingNumber"`
	LanguageId    string `json:"languageId"`
	Signature     string `json:"signature"`
}
