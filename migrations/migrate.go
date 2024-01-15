package migrations

import (
	"log"

	"www.github.com/gokhankocer/ticket-api/database"
	"www.github.com/gokhankocer/ticket-api/entities"
)

func Migrate() {
	err := database.DB.AutoMigrate(&entities.TicketOption{}, &entities.Ticket{})
	if err != nil {
		log.Fatalf("Migration hatası: %v", err)
	}
}
