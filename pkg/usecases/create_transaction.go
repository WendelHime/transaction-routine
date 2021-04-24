package usecases

import (
	"sync"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
	"github.com/WendelHime/transaction-routine/pkg/domain/services"
)

type createTrans struct {
	tranServ          *services.TransactionService
	accountServ       *services.AccountService
	accountsInProcess map[int]bool
	mutex             *sync.Mutex
}

// NewUCTransactionCreator builds a new implementation of UCTransactionCreator
func NewUCTransactionCreator(tranServ *services.TransactionService, accountServ *services.AccountService) UCTransactionCreator {
	return &createTrans{tranServ: tranServ, accountServ: accountServ, accountsInProcess: make(map[int]bool), mutex: new(sync.Mutex)}
}

// CreateTransaction register a transaction
func (u *createTrans) CreateTransaction(transaction *models.Transaction) error {

	acc, err := u.accountServ.Get(transaction.Account.ID)
	if err != nil {
		return err
	}

	newCredit := acc.AvailableCreditLimit + transaction.Amount
	if newCredit < 0 {
		return models.ErrAmountBeyondCreditLimit
	}

	err = u.tranServ.Create(transaction)
	if err != nil {
		return err
	}

	acc.AvailableCreditLimit = newCredit
	return u.accountServ.UpdateCreditLimit(acc)
}
