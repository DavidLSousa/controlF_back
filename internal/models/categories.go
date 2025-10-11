package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(100);not null"`
	Icon      string `gorm:"type:varchar(10)"` // Para armazenar emojis como '📚', '💪🏼'
}
