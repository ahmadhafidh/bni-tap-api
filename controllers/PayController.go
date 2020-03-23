package controllers

import (
	"BNI-TAP/commonsetting"
	"BNI-TAP/dbhelper"
	"BNI-TAP/dbmodel"
	"BNI-TAP/helper"
	"BNI-TAP/model"
	"encoding/json"
	"log"
	"net/http"
)

func Payment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// Api Model
	var ReqClient model.ReqClient
	var ResClient model.ResClient
	var ReqPayBNI model.ReqPayBNI
	var SignedReqPayBNI model.SignedReqPayBNI
	var ResPayBNI model.ResPayBNI
	var ResBniFail model.ResBNIFail
	var ResInqBNI model.ResInqBni
	var responsebni string
	var rc int
	var trxid string
	// End Of Api Model
	// Db Model
	var Vabilling dbmodel.Vabilling
	// End Of Db Model
	var result interface{}
	var URL = commonsetting.URL + commonsetting.PAYMENT + "&token=DwTkE0HZD6it4xwcBwiKWSG9n716wu7MV8IkPSfvRbAS0AtgPXhz9b"
	errors := decoder.Decode(&ReqClient)
	if errors != nil {
		log.Print("errors")
	}
	Vabilling = dbhelper.GetInqVabilling(ReqClient.TRXID, "0")
	resinquiry := json.Unmarshal([]byte(Vabilling.Inquirydata), &ResInqBNI)
	if resinquiry != nil {
		log.Print(resinquiry)
	}

	dJson, _ := json.Marshal(ReqPayBNI)
	jwt := helper.GenSignature(dJson)
	SignedReqPayBNI = helper.SignReqPay(ReqPayBNI, jwt)
	requestjson, _ := json.Marshal(SignedReqPayBNI)
	responsebni = helper.CallApiBNI(requestjson, URL)
	parsuccess := json.Unmarshal([]byte(responsebni), &ResPayBNI)
	if parsuccess != nil {
		rc = 500
		result = "Errors"
	} else {
		if ResPayBNI.BillingNumber == "" {
			parserror := json.Unmarshal([]byte(responsebni), &ResBniFail)
			log.Println(parserror)
			rc = 500
			result = ResBniFail
		} else {
			rc = 200
			trxid = ReqClient.TRXID
			result = ResPayBNI
		}
	}
	ResClient.Status = rc
	ResClient.TRXID = trxid
	ResClient.Data = result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResClient)
}
