package http

import (
	"time"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

// Account represents an user account
type Account struct {
	ID                   int     `json:"account_id"`
	DocumentNumber       string  `json:"document_number"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}

// Transaction represents a transaction executed by the user related to an
// account
type Transaction struct {
	ID        int                  `json:"transaction_id"`
	AccountID int                  `json:"account_id"`
	Operation models.OperationType `json:"operation_type_id"`
	Amount    float64              `json:"amount"`
	EventDate time.Time            `json:"event_date"`
}
