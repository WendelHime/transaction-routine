package usecases

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
	"github.com/WendelHime/transaction-routine/pkg/domain/services"
)

type accountGetter struct {
	accServ *services.AccountService
}

// NewUCAccountGetter builds a new implementation of account getter
func NewUCAccountGetter(accServ *services.AccountService) UCAccountGetter {
	return &accountGetter{accServ: accServ}
}

// GetAccount search by account with supplied identifier
func (u *accountGetter) GetAccount(accID int) (*models.Account, error) {
	return u.accServ.Get(accID)
}
