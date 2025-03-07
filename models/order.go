package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
    ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
    ProductID uuid.UUID  `json:"product_id"`
    Product   Product `json:"product" gorm:"foreignKey:ProductID"`
    UserID    uuid.UUID  `json:"user_id"`
    User      User    `json:"user" gorm:"foreignKey:UserID"`
}