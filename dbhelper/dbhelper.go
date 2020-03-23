package dbhelper

import (
	"BNI-TAP/dbmodel"
)

func GetInqVabilling(transactionid string, status string) dbmodel.Vabilling {
	var Vabilling dbmodel.Vabilling
	db, _ := ConnectDbMySql()
	db.Table(Vabilling.TableName()).Where("transaction_id = ? AND status = ?", transactionid, status).Scan(&Vabilling)
	return Vabilling
}
func Insertlog() {

}
func GetCredentials() {

}
