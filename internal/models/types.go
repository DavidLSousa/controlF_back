package models

type UserType uint8

const (
	UserTypePersonal      UserType = 1 // Usuário de conta pessoal
	UserTypeCompanyAdmin  UserType = 2 // Usuário admin de uma empresa
	UserTypeCompanyMember UserType = 3 // Usuário membro de uma empresa
)

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "INCOME"  // Entrada
	TransactionTypeExpense TransactionType = "EXPENSE" // Saída
)

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
)
