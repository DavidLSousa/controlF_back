package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name   string `gorm:"type:varchar(100);not null"`
	Icon   string `gorm:"type:varchar(10)"` // Para armazenar emojis como '📚', '💪🏼'
	Color  string `gorm:"type:varchar(7)"`  // Ex: #FFFFFF
	Status Status `gorm:"type:varchar(10);not null"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	User   *User     `gorm:"foreignKey:UserID"`

	Transactions    []Transaction  `gorm:"foreignKey:CategoryID"`
	PaymentMethodID *uuid.UUID     // Pode ser nulo se a categoria não estiver vinculada a um método de pagamento específico
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
}
