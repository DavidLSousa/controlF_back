package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Installment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Current       int          `gorm:"default:0"`
	Total         int          `gorm:"default:0"`
	TransactionID uuid.UUID    `gorm:"type:uuid;not null;uniqueIndex"`
	Transaction   *Transaction `gorm:"foreignKey:TransactionID"`
}
