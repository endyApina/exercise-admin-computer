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

type UpdateComputerAllocationRequestBody struct {
	ComputerID           string `json:"computer_id"`
	EmployeeAbbreviation string `json:"employee_abbreviation"`
}
