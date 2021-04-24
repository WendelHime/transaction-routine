package services

import (
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

type accRepo struct{}

// Create register an account on persistence service
func (u *accRepo) Create(account *models.Account) error {
	return nil
}

// Get searches by account using the supplied identifier and returns
// a account reference from persistence service
func (u *accRepo) Get(accountID int) (*models.Account, error) {
	return nil, nil
}

// UpdateCreditLimit update credit limit from account on persistence service
func (u *accRepo) UpdateCreditLimit(account *models.Account) error {
	return nil
}

func TestCreateAccount(t *testing.T) {
	serv := NewAccountService(new(accRepo), new(accRepo), new(accRepo))
	var tests = []struct {
		name     string
		expected error
		given    *models.Account
	}{
		{"empty document", models.ErrMissingRequiredField, new(models.Account)},
		{"success", nil, &models.Account{DocumentNumber: "99999999999"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := serv.Create(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestGetAccount(t *testing.T) {
	serv := NewAccountService(new(accRepo), new(accRepo), new(accRepo))
	var tests = []struct {
		name     string
		expected error
		given    int
	}{
		{"invalid account ID", models.ErrMissingRequiredField, 0},
		{"success", nil, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := serv.Get(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestUpdateCreditLimit(t *testing.T) {
	serv := NewAccountService(new(accRepo), new(accRepo), new(accRepo))
	var tests = []struct {
		name     string
		expected error
		given    *models.Account
	}{
		{"invalid account id", models.ErrMissingRequiredField, new(models.Account)},
		{"success", nil, &models.Account{ID: 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := serv.UpdateCreditLimit(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
