// Package models contain all data transfer objects used in this application
// for creating standards of communication between all the layers.
package models

import "time"

// Account represents an user account
type Account struct {
	ID             int
	DocumentNumber string
}

// OperationType represents a enum type for operations
type OperationType int

const (
	// OperationTypeCompraVista used when the user buy something and will pay
	// in cash
	OperationTypeCompraVista OperationType = iota + 1
	// OperationTypeCompraParcelada used when the user buy something in parcels
	OperationTypeCompraParcelada
	// OperationTypeSaque used when the user extract some money from account
	OperationTypeSaque
	// OperationTypePagamento used when the user is paying his bills from
	// account
	OperationTypePagamento
)

// Transaction represents a transaction executed by the user related to an
// account
type Transaction struct {
	ID        int
	Account   Account
	Operation OperationType
	Amount    float64
	EventDate time.Time
}
