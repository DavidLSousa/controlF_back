package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name   string    `gorm:"type:varchar(100);not null"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	User   *User     `gorm:"foreignKey:UserID"`

	Categories   []Category    `gorm:"foreignKey:PaymentMethodID"`
	Transactions []Transaction `gorm:"foreignKey:PaymentMethodID"`
}
