package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	

	ID uuid.UUID `json:"ID" gorm:"primaryKey"`

	  
	 
	Name string `json:"name"`
	
}