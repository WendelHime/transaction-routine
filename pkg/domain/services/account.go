// Package services contain all interactions with external libraries.
// This package also validate args before interacting with any implementation.
package services

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/interfaces"
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

// AccountService interact with implementations validating args before
// applying
type AccountService struct {
	creator interfaces.AccountCreator
	getter  interfaces.AccountGetter
}

// NewAccountService build a new account service
func NewAccountService(creator interfaces.AccountCreator, getter interfaces.AccountGetter) *AccountService {
	return &AccountService{creator: creator, getter: getter}
}

// Create register an account on persistence if supplied account is valid.
// If account have a invalid DocumentNumber, raises models.ErrMissingRequiredField
func (s *AccountService) Create(account *models.Account) error {
	if account.DocumentNumber == "" {
		return models.ErrMissingRequiredField
	}
	return s.creator.Create(account)
}

// Get searches by an account through identifier.
// If accountID is invalid, raises models.ErrMissingRequiredField
func (s *AccountService) Get(accountID int) (*models.Account, error) {
	if accountID <= 0 {
		return nil, models.ErrMissingRequiredField
	}

	return s.getter.Get(accountID)
}
