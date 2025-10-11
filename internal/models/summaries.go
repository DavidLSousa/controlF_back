package models

import (
	"time"

	"github.com/google/uuid"
)

type Summary struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	User      *User

	Month    string  `gorm:"type:varchar(20);not null"` // Ex: "Dezembro"
	TotalIn  float64 `gorm:"not null"`
	TotalOut float64 `gorm:"not null"`
	Balance  float64 `gorm:"not null"`
}
