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

	Description     string          `gorm:"type:varchar(255);not null"`
	Amount          decimal.Decimal `gorm:"type:decimal(10,2);not null"` // Ex: 12345678.99
	Date            time.Time       `gorm:"not null"`                    // Data da transação/vencimento
	Type            TransactionType `gorm:"type:varchar(10);not null"`
	IsRecurring     bool            `gorm:"default:false"`    // Para contas recorrentes (Academia)
	IsPaid          bool            `gorm:"default:false"`    // Para marcar se já foi paga
	IsInstallment   bool            `gorm:"default:false"`    // Se é uma parcela (Flip 5)
	InstallmentInfo string          `gorm:"type:varchar(20)"` // Opcional: para guardar "3/12"

	// Relacionamentos
	CategoryID uuid.UUID `gorm:"not null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
}
