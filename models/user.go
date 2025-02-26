package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name string `json:"name"`
	
}