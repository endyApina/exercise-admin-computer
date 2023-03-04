package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/endyApina/exercise-admin-computer/db"
	"github.com/endyApina/exercise-admin-computer/lib/idgenerator"
	"github.com/endyApina/exercise-admin-computer/types/models"
	"github.com/go-chi/chi"
)

type Handler struct {
	store       db.DataStore
	idGenerator idgenerator.IdGenerator
}

func NewHttpHandler(store db.DataStore, idGenerator idgenerator.IdGenerator) *Handler {
	return &Handler{
		store:       store,
		idGenerator: idGenerator,
	}
}

func (handler *Handler) TestHealth(w http.ResponseWriter, r *http.Request) {
	message := "status ok"
	handler.sendResponse(w, http.StatusOK, message)
}

func (handler *Handler) CreateComputer(w http.ResponseWriter, r *http.Request) {
	//validate json header
	err := handler.validateHeader(w, r)
	if err != nil {
		log.Println("error validating header")
		return
	}

	//parse data
	var newComputer *models.Computer
	newComputer.ComputerID = handler.idGenerator.Generate() //generate unique id for computer

	err = handler.parseData(r, newComputer)
	if err != nil {
		errorMessage := errors.New("could not parse data")
		handler.sendResponse(w, http.StatusBadRequest, errorMessage)
		return
	}

	//store data in db
	err = handler.store.CreateComputer(newComputer)
	if err != nil {
		errorMessage := errors.New("error saving data in db")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	//validate no of computers for this user
	allComputers, err := handler.store.GetComputersByEmployeeName(newComputer.EmployeeAbbreviation)
	if err != nil {
		errorMessage := errors.New("error getting all customer computer information")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	if len(allComputers) >= 3 {
		go handler.handleNotifyExcessComputerAllocation(allComputers)
	}

	handler.sendResponse(w, http.StatusOK, "computer created successfully")
}

func (handler *Handler) GetAllComputers(w http.ResponseWriter, r *http.Request) {
	//validate json header
	err := handler.validateHeader(w, r)
	if err != nil {
		log.Println("error validating header")
		return
	}

	//get all computers in db
	allComputers, err := handler.store.GetAllComputers()
	if err != nil {
		errorMessage := errors.New("error getting all computer")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	handler.sendResponse(w, http.StatusOK, allComputers)
}

func (handler *Handler) GetComputersByEmployeeName(w http.ResponseWriter, r *http.Request) {
	//validate json header
	err := handler.validateHeader(w, r)
	if err != nil {
		log.Println("error validating header")
		return
	}

	employeeName := chi.URLParam(r, "employee_abbreviation")
	allComputers, err := handler.store.GetComputersByEmployeeName(employeeName)
	if err != nil {
		errorMessage := errors.New("error getting all customer computer information")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	handler.sendResponse(w, http.StatusOK, allComputers)
}

func (handler *Handler) GetComputerByComputerID(w http.ResponseWriter, r *http.Request) {
	//validate json header
	err := handler.validateHeader(w, r)
	if err != nil {
		log.Println("error validating header")
		return
	}

	computerID := chi.URLParam(r, "computer_id")
	allComputers, err := handler.store.GetComputerByID(computerID)
	if err != nil {
		errorMessage := errors.New("error getting all computer information")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	handler.sendResponse(w, http.StatusOK, allComputers)
}

func (handler *Handler) DeleteComputerByComputerID(w http.ResponseWriter, r *http.Request) {
	//validate json header
	err := handler.validateHeader(w, r)
	if err != nil {
		log.Println("error validating header")
		return
	}

	computerID := chi.URLParam(r, "computer_id")
	allComputers, err := handler.store.DeleteComputer(computerID)
	if err != nil {
		errorMessage := errors.New("error getting all computer information")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	handler.sendResponse(w, http.StatusOK, allComputers)
}

func (handler *Handler) UpdateComputerAllocation(w http.ResponseWriter, r *http.Request) {
	//validate json header
	err := handler.validateHeader(w, r)
	if err != nil {
		log.Println("error validating header")
		return
	}

	//parse data
	var requestBody *models.UpdateComputerAllocationRequestBody

	err = handler.parseData(r, requestBody)
	if err != nil {
		errorMessage := errors.New("could not parse data")
		handler.sendResponse(w, http.StatusBadRequest, errorMessage)
		return
	}

	//validate no of computers for this user
	err = handler.store.UpdateComputerAllocation(requestBody.ComputerID, requestBody.EmployeeAbbreviation)
	if err != nil {
		errorMessage := errors.New("error getting all customer computer information")
		handler.sendResponse(w, http.StatusInternalServerError, errorMessage)
		return
	}

	handler.sendResponse(w, http.StatusOK, "computer allocation update successfully")
}

func (handler *Handler) validateHeader(w http.ResponseWriter, r *http.Request) error {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorMessage := errors.New("content type must be application/json")
		handler.sendResponse(w, http.StatusBadRequest, errorMessage)
		return errorMessage
	}
	return nil
}

func (handler *Handler) parseData(request *http.Request, model interface{}) error {
	return json.NewDecoder(request.Body).Decode(model)
}

func (handler *Handler) sendResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body == nil {
		return
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println("Could not parse body", err)
	}
}

func (handler *Handler) handleNotifyExcessComputerAllocation(computerData []*models.Computer) {
	//not implemented

}
