// Package interfaces contain specifications for interaction with external libraries
package interfaces

import "github.com/WendelHime/transaction-routine/pkg/domain/models"

// AccountCreator represents an interface for account creation.
type AccountCreator interface {
	// Create register an account on persistence service
	Create(account *models.Account) error
}

// AccountGetter represents an interface for getting operations from account
type AccountGetter interface {
	// Get searches by account using the supplied identifier and returns
	// a account reference from persistence service
	Get(accountID int) (*models.Account, error)
}

// TransactionCreator represents an interface for creation of transactions
type TransactionCreator interface {
	// Create register a transaction on persistence service
	Create(transaction *models.Transaction) error
}
