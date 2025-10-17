package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Description string          `gorm:"type:varchar(255);not null"`
	Amount      decimal.Decimal `gorm:"type:decimal(10,2);not null"` // Ex: 12345678.99
	Date        time.Time       `gorm:"not null"`                    // Data da transação/vencimento
	Type        TransactionType `gorm:"type:varchar(10);not null"`
	IsRecurring bool            `gorm:"default:false"` // Para contas recorrentes (Academia)
	IsPaid      bool            `gorm:"default:false"` // Para marcar se já foi paga

	// Relacionamentos
	Installment     *Installment
	CategoryID      uuid.UUID      `gorm:"type:uuid;not null"`
	Category        *Category      `gorm:"foreignKey:CategoryID"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null;index"`
	User            *User          `gorm:"foreignKey:UserID"`
	PaymentMethodID uuid.UUID      `gorm:"type:uuid;not null"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	CompanyID       *uuid.UUID     `gorm:"type:uuid"`
	Company         *Company       `gorm:"foreignKey:CompanyID"`
}
