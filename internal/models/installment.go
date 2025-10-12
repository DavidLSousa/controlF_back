package models

import "github.com/google/uuid"

type Installment struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Current int       `gorm:"default:0"` // Parcela atual (ex: 1 de 12)
	Total   int       `gorm:"default:0"` // Total de parcelas (ex: 12)

	TransactionID uuid.UUID   `gorm:"type:uuid;not null"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
}
