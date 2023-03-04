package models

// Computer holds the required attributes of the computer object
type Computer struct {
	MacAddress           string `json:"mac_address"`
	ComputerName         string `json:"computer_name"`
	IPAddress            string `json:"ip_address"`
	EmployeeAbbreviation string `json:"employee_abbreviation"`
	Description          string `json:"description"`
}
