package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/endyApina/exercise-admin-computer/db"
)

type Handler struct {
	store db.DataStore
}

func NewHttpHandler(store db.DataStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (handler *Handler) TestHealth(w http.ResponseWriter, r *http.Request) {
	message := "status ok"
	SendResponse(w, http.StatusOK, message)
}

func SendResponse(w http.ResponseWriter, statusCode int, body interface{}) {
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
