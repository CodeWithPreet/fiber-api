package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	 
	Name string `json:"name"`
	SerialNo string `json:"serial_no"`
}