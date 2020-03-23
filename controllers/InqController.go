package controllers

import (
	"BNI-TAP/commonsetting"
	"BNI-TAP/helper"
	"BNI-TAP/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Inquiry(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ReqClient model.ReqClient
	var ResClient model.ResClient
	var ReqInqBNI model.ReqInqBNI
	var SignedReqInqBNI model.SignedReqInqBNI
	var ResInqBNI model.ResInqBni
	var ResBniFail model.ResBNIFail
	var responsebni string
	var rc int
	var trxid string
	var result interface{}
	var URL = commonsetting.URL + commonsetting.INQUIRY + "&token=DwTkE0HZD6it4xwcBwiKWSG9n716wu7MV8IkPSfvRbAS0AtgPXhz9b"

	errors := decoder.Decode(&ReqClient)
	if errors != nil {
		log.Print("errors")
	}
	ReqInqBNI.BillingNumber = ReqClient.BillingNumber
	ReqInqBNI.LanguageId = "ID"

	dJson, _ := json.Marshal(ReqInqBNI)

	jwt := helper.GenSignature(dJson)

	SignedReqInqBNI.BillingNumber = ReqInqBNI.BillingNumber
	SignedReqInqBNI.LanguageId = ReqInqBNI.LanguageId
	SignedReqInqBNI.Signature = jwt

	requestjson, _ := json.Marshal(SignedReqInqBNI)
	responsebni = helper.CallApiBNI(requestjson, URL)
	parsuccess := json.Unmarshal([]byte(responsebni), &ResInqBNI)
	if parsuccess != nil {
		rc = 500
		result = "Errors"
	} else {
		if ResInqBNI.BillingNumber == "" {
			parserror := json.Unmarshal([]byte(responsebni), &ResBniFail)
			log.Println(parserror)
			rc = 500
			result = ResBniFail
		} else {
			rc = 200
			result = ResInqBNI
			trxid = time.Now().Format("YmdHis")

		}
	}
	ResClient.Status = rc
	ResClient.TRXID = trxid
	ResClient.Data = result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResClient)
}
