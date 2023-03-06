package models

import "time"

type BaseModel struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// Computer holds the required attributes of the computer object
type Computer struct {
	BaseModel
	ComputerID           string `json:"computer_id"`
	MacAddress           string `json:"mac_address"`
	ComputerName         string `json:"computer_name"`
	IPAddress            string `json:"ip_address"`
	EmployeeAbbreviation string `json:"employee_abbreviation"`
	Description          string `json:"description"`
}

type CreateComputerRequest struct {
	MacAddress           string `json:"mac_address"`           //required
	ComputerName         string `json:"computer_name"`         //required
	IPAddress            string `json:"ip_address"`            //required
	EmployeeAbbreviation string `json:"employee_abbreviation"` //optional
	Description          string `json:"description"`           //optional
}

type UpdateComputerAllocationRequestBody struct {
	ComputerID           string `json:"computer_id"`
	EmployeeAbbreviation string `json:"employee_abbreviation"`
}

type NotificationRequestBody struct {
	Level                string `json:"level"`
	EmployeeAbbreviation string `json:"employeeAbbreviation"`
	Message              string `json:"message"`
}

type HttpResponseBody struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}
