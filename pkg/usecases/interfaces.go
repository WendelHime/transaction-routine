// Package usecases contain use case definitions and implementations
package usecases

import "github.com/WendelHime/transaction-routine/pkg/domain/models"

// UCAccountCreator specify what a usecase for creating accounts must have
type UCAccountCreator interface {
	// CreateAccount register an account
	CreateAccount(acc *models.Account) error
}

// UCAccountGetter specify which operations account getter usecase must have
type UCAccountGetter interface {
	// GetAccount search by account with supplied identifier
	GetAccount(accID int) (*models.Account, error)
}

// UCTransactionCreator specify which operations a transaction creator must
// have
type UCTransactionCreator interface {
	// CreateTransaction register a transaction
	CreateTransaction(transaction *models.Transaction) error
}
