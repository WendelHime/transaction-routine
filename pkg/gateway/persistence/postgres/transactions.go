package postgres

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/interfaces"
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

type transactionRepository struct {
	conn *Connection
}

// NewTransactionCreator builds a new transaction creator
func NewTransactionCreator(conn *Connection) interfaces.TransactionCreator {
	return &transactionRepository{conn: conn}
}

// Create register a transaction on persistence service
func (r *transactionRepository) Create(transaction *models.Transaction) error {
	db, err := r.conn.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(
		`INSERT INTO transactions(account_id, operationtype_id, amount)
		VALUES ($1, $2, $3) RETURNING id, event_date `,
		transaction.Account.ID,
		transaction.Operation,
		transaction.Amount).Scan(&transaction.ID, &transaction.EventDate)
	if err != nil {
		return err
	}

	return nil
}
