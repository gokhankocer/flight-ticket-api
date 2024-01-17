package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"www.github.com/gokhankocer/ticket-api/database"
	"www.github.com/gokhankocer/ticket-api/entities"
	"www.github.com/gokhankocer/ticket-api/models"
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
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var ticket entities.Ticket
	result := database.DB.Preload("TicketOption").First(&ticket, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)

}

func PurchaseTicket(w http.ResponseWriter, r *http.Request) {
	// URL'den ticket_option ID'sini al
	vars := mux.Vars(r)
	optionID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid Ticket Option ID format", http.StatusBadRequest)
		return
	}

	var ticketRequest models.PurchaseRequest

	err = json.NewDecoder(r.Body).Decode(&ticketRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var option entities.TicketOption
	result := database.DB.First(&option, optionID)
	if result.Error != nil {
		http.Error(w, "Ticket Option not found", http.StatusNotFound)
		return
	}

	if option.Allocation < ticketRequest.Quantity {
		http.Error(w, "Not enough tickets available", http.StatusBadRequest)
		return
	}
	for i := 0; i < ticketRequest.Quantity; i++ {
		ticket := entities.Ticket{
			TicketOptionID: uint(optionID),
			PurchaseDate:   time.Now(),
		}
		database.DB.Create(&ticket)
	}
	option.Allocation -= ticketRequest.Quantity
	database.DB.Save(&option)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Tickets successfully purchased")
}
