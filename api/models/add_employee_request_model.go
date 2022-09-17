package models

type AddEmployeeRequestModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
