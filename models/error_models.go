package models

type ErrorResponse struct {
	Status_Code int `json:"status_code"`
	Message string	`json:"message"`
}