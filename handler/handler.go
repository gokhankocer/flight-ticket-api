package handler

import (
	"encoding/json"
	"net/http"

	"www.github.com/gokhankocer/ticket-api/database"
	"www.github.com/gokhankocer/ticket-api/entities"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var newTicketOption entities.TicketOption

	err := json.NewDecoder(r.Body).Decode(&newTicketOption)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := database.DB.Create(&newTicketOption)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTicketOption)
}

func GetTicektByID(w http.ResponseWriter, r *http.Request) {
	// Bilet bilgisini almak için işlemler
}

func PurchaseTicket(w http.ResponseWriter, r *http.Request) {
	// Bilet seçeneğinden satın alma işlemleri
}
