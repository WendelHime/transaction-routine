package services

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/interfaces"
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

// TransactionService interact with transaction implementations and validate
// args before performing the operation
type TransactionService struct {
	creator interfaces.TransactionCreator
}

// NewTransactionService builds a new transaction service
func NewTransactionService(creator interfaces.TransactionCreator) *TransactionService {
	return &TransactionService{creator: creator}
}

// Create register a new transaction on persistence service if transaction is valid.
// If supplied account id is invalid, it will raise an error models.ErrMissingRequiredField
// If supplied amount is incompatible with the operation, it will raise an error models.ErrInvalidAmount
func (s *TransactionService) Create(transaction *models.Transaction) error {
	if transaction.Account.ID <= 0 {
		return models.ErrMissingRequiredField
	}

	switch transaction.Operation {
	case models.OperationTypeCompraVista, models.OperationTypeCompraParcelada, models.OperationTypeSaque:
		if transaction.Amount > 0 {
			return models.ErrInvalidAmount
		}
	case models.OperationTypePagamento:
		if transaction.Amount <= 0 {
			return models.ErrInvalidAmount
		}
	}

	return s.creator.Create(transaction)
}
