package model

type ResClient struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	TRXID  string      `json:"trx_id"`
}

// type DataResInqClient struct {
// 	Signature string `json:"signature"`
// }
