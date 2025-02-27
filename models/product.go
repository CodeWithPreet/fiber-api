package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	 
	Name string `json:"name"`
	SerialNo string `json:"serial_no"`
}