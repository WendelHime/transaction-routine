package usecases

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
	"github.com/WendelHime/transaction-routine/pkg/domain/services"
)

type accountCreator struct {
	accServ *services.AccountService
}

// NewUCAccountCreator builds a new implementation of account creator
func NewUCAccountCreator(accServ *services.AccountService) UCAccountCreator {
	return &accountCreator{accServ: accServ}
}

// CreateAccount register an account
func (u *accountCreator) CreateAccount(acc *models.Account) error {
	return u.accServ.Create(acc)
}
