package models

type Response struct {
	Status           bool        `json:"Status"`
	ErrorCode        string      `json:"ErrorCode"`
	ErrorDescription string      `json:"ErrorDescription"`
	Data             interface{} `json:"Data"`
}
