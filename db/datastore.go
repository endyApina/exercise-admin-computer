package db

import "github.com/endyApina/exercise-admin-computer/types/models"

type DataStore interface {
	ComputerStore
}

type ComputerStore interface {
	CreateComputer(computerData *models.Computer) error
	GetComputerByID(computerId string) (*models.Computer, error)
	GetComputerByEmployeeName(employeeName string) (*models.Computer, error)
	GetAllComputers() ([]*models.Computer, error)
	UpdateComputer(computerData *models.Computer) error
	DeleteComputer(cumputerId string) (*models.Computer, error)
	UnAssignComputer(computerId, employeeId string) error
	UpdateComputerAllocation(computerId, employeeId string) error
}
