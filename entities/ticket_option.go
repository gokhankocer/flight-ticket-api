package entities

import "gorm.io/gorm"

type TicketOption struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null"`
	Allocation int    `gorm:"type:int;not null"` // Toplam bilet sayısını temsil eder.
}
