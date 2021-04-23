package usecases

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
	"github.com/WendelHime/transaction-routine/pkg/domain/services"
)

type createTrans struct {
	tranServ *services.TransactionService
}

// NewUCTransactionCreator builds a new implementation of UCTransactionCreator
func NewUCTransactionCreator(tranServ *services.TransactionService) UCTransactionCreator {
	return &createTrans{tranServ: tranServ}
}

// CreateTransaction register a transaction
func (u *createTrans) CreateTransaction(transaction *models.Transaction) error {
	return u.tranServ.Create(transaction)
}
