package model

type ResBNIFail struct {
	Response Response `json:"response"`
}
type Response struct {
	Parameters Parameters `json:"parameters"`
}
type Parameters struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ErrorMessage    string `json:"errorMessage"`
}
