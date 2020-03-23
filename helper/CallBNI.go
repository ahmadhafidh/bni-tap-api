package helper

import (
	"BNI-TAP/commonsetting"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// CallApiBNI function for calling BNI-TAP API
func CallApiBNI(requestbni []byte, URL string) string {
	var result string
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(requestbni))
	request.Header.Set("Content-Type", commonsetting.ContentType)
	request.Header.Set("x-api-key", "f2666c76-365c-4dce-8bd3-4e3910a8f30a")
	res, err := client.Do(request)
	if err != nil {
		result = err.Error()
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		result = err.Error()
	} else {
		result = string(body)
	}
	return result
}
func GetTokenBNi(requestbni []byte, URL string) string {
	var result string
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(requestbni))
	request.Header.Set("Content-Type", commonsetting.ContentType)
	request.Header.Set("x-api-key", "f2666c76-365c-4dce-8bd3-4e3910a8f30a")
	res, err := client.Do(request)
	if err != nil {
		result = err.Error()
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		result = err.Error()
	} else {
		result = string(body)
	}
	return result
}
