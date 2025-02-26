package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name string `json:"name"`
	SerialNo string `json:"serial_no"`
}