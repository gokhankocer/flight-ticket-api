package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"www.github.com/gokhankocer/ticket-api/handler"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/ticket_options", handler.CreateTicket).Methods("POST")
	r.HandleFunc("/ticket/{id}", handler.GetTicektByID).Methods("GET")
	r.HandleFunc("/ticket_options/{id}/purchases", handler.PurchaseTicket).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
