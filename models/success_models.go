package models

type SuccessResponse struct{
	Status_Code int 	`json:"status_code"`
	Message string 		`json:"message"`
	Data any			`json:"data"`
}