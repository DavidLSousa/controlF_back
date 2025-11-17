package models

// users

// transactions
type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "INCOME"  // Entrada
	TransactionTypeExpense TransactionType = "EXPENSE" // Saída
)

// categories
type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
)

// claims
type Claims string

const (
	ClaimLimited       Claims = "LIMITED"
	ClaimVerified      Claims = "VERIFIED" // email
	ClaimComplete      Claims = "COMPLETE" // cpf, end, phone
	ClaimAdministrator Claims = "ADMIN"
	ClaimBetaTester    Claims = "BETATESTER"
)
