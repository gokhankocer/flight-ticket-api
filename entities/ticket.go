package entities

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketOptionID uint         `gorm:"not null"`
	TicketOption   TicketOption `gorm:"foreignKey:TicketOptionID"` // TicketOption ile ilişki
	PurchaseDate   time.Time    `gorm:"type:time;not null"`
}
