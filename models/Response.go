package models

type Response struct {
	Value   string `json:"value"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
