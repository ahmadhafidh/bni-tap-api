package model

type ReqClient struct {
	BillingNumber string `json:"billingNumber"`
	TRXID         string `json:"trx_id"`
}
