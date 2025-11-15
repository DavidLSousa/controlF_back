package models

// users
type UserType uint8

const (
	UserTypePersonal      UserType = 1 // Usuário de conta pessoal
	UserTypeCompanyAdmin  UserType = 2 // Usuário admin de uma empresa
	UserTypeCompanyMember UserType = 3 // Usuário membro de uma empresa
)

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
