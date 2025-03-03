package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	// CreatedAt time.Time
	ProductRefer uuid.UUID `json:"product_id"`
	Product Product `gorm:"foreignKey:ProductRefer"`
	UserRefer uuid.UUID `json:"user_id"`
	User User `gorm:"foreignKey:UserRefer"`
}