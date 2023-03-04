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
	err := store.postgresClient.Create(computerData).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (store *postgresStore) GetComputerByID(computerId string) (*models.Computer, error) {
	var computer models.Computer
	err := store.postgresClient.Where("computer_id = ?", computerId).Find(&computer).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &computer, nil
}

func (store *postgresStore) GetComputersByEmployeeName(employeeName string) ([]*models.Computer, error) {
	var allComputer []*models.Computer
	err := store.postgresClient.Where("employee_abbreviation = ?", employeeName).Find(allComputer).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return allComputer, nil
}

func (store *postgresStore) GetAllComputers() ([]*models.Computer, error) {
	var allComputers []models.Computer
	err := store.postgresClient.Find(allComputers).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return []*models.Computer{}, nil
}

func (store *postgresStore) UpdateComputer(computerData *models.Computer) error {
	err := store.postgresClient.Model(models.Computer{}).Where("computer_id = ?", computerData.ComputerID).Updates(computerData).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (store *postgresStore) DeleteComputer(computerId string) (*models.Computer, error) {
	err := store.postgresClient.Where("computer_id = ?", computerId).Delete(models.Computer{}).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.Computer{}, nil
}

func (store *postgresStore) UnAssignComputer(computerId, employeeAbbreviation string) error {
	err := store.postgresClient.Where("computer_id = ? AND employee_abbreviation = ?", computerId, employeeAbbreviation).Delete(models.Computer{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (store *postgresStore) UpdateComputerAllocation(computerId, employeeAbbreviation string) error {
	err := store.postgresClient.Model(models.Computer{}).Where("computer_id = ?", computerId).Update("employee_abbreviation", employeeAbbreviation).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
