package postgres

import (
	"fmt"
	"log"

	"github.com/endyApina/exercise-admin-computer/db"

	"github.com/endyApina/exercise-admin-computer/config"
	"github.com/endyApina/exercise-admin-computer/types/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(secrets *config.Secrets) (db.DataStore, error) {
	url := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=%s", secrets.DatabaseHost, secrets.DatabasePassword, secrets.DatabaseName, secrets.DatabasePort)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.AutoMigrate(models.Computer{})
	return &postgresStore{
		postgresClient: db,
	}, nil
}

var _ db.ComputerStore = &postgresStore{}

type postgresStore struct {
	postgresClient *gorm.DB
}

func (store *postgresStore) CreateComputer(computerData *models.Computer) error {
	err := store.postgresClient.Create(computerData)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (store *postgresStore) GetComputerByID(computerId string) (*models.Computer, error) {
	//NOT IMPLEMENTED
	return &models.Computer{}, nil
}

func (store *postgresStore) GetComputerByEmployeeName(employeeName string) (*models.Computer, error) {
	//NOT IMPLEMENTED
	return &models.Computer{}, nil
}

func (store *postgresStore) GetAllComputers() ([]*models.Computer, error) {
	//NOT IMPLEMENTED
	return []*models.Computer{}, nil
}

func (store *postgresStore) UpdateComputer(computerData *models.Computer) error {
	//NOT IMPLEMENTED
	return nil
}

func (store *postgresStore) DeleteComputer(cumputerId string) (*models.Computer, error) {
	//NOT IMPLEMENTED
	return &models.Computer{}, nil
}

func (store *postgresStore) UnAssignComputer(computerId, employeeId string) error {
	//NOT IMPLEMENTED
	return nil
}

func (store *postgresStore) UpdateComputerAllocation(computerId, employeeId string) error {
	//NOT IMPLEMENTED
	return nil
}
