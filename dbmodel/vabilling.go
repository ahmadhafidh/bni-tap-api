package dbmodel

type Vabilling struct {
	Idmerchant     string `gorm:"column:id_merchant"`
	Merchantid     string `gorm:"column:merchant_id"`
	Merchantname   string `gorm:"column:merchant_name"`
	Billingnumber  string `gorm:"column:billing_number"`
	Transactionid  string `gorm:"column:transaction_id"`
	Billamount     string `gorm:"column:bill_amount"`
	Feebni         string `gorm:"column:fee_bni"`
	Feespe         string `gorm:"column:fee_spe"`
	Feetransaction string `gorm:"column:fee_transaction"`
	Feetotal       string `gorm:"column:fee_total"`
	Journalnumber  string `gorm:"column:journal_number"`
	Mpaymentdate   string `gorm:"column:m_payment_date"`
	Paymentdate    string `gorm:"column:payment_date"`
	Status         string `gorm:"column:status"`
	Retrycount     string `gorm:"column:retry_count"`
	Reqpayparam    string `gorm:"column:req_pay_param"`
	Inquirydata    string `gorm:"column:inquiry_data"`
	Paymentdata    string `gorm:"column:payment_data"`
}

func (Vabilling) TableName() string {
	return "vabilling"
}
