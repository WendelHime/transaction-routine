package postgres

import (
	"database/sql"
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

func TestCreateAccount(t *testing.T) {
	creator := NewAccountCreator(NewConnection("127.0.0.1", 5432, "docker", "docker", "routine"))
	var tests = []struct {
		name     string
		expected error
		given    *models.Account
	}{
		{"success", nil, &models.Account{DocumentNumber: "99999999900", AvailableCreditLimit: 10}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := creator.Create(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestGetAccount(t *testing.T) {
	getter := NewAccountGetter(NewConnection("127.0.0.1", 5432, "docker", "docker", "routine"))
	var tests = []struct {
		name     string
		expected error
		given    int
	}{
		{"invalid id", sql.ErrNoRows, 0},
		{"success", nil, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			account, actual := getter.Get(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %s, actual %s", tt.given, tt.expected, actual)
			}

			if account.ID != tt.given {
				t.Errorf("(%d): expected %d, actual %d", tt.given, tt.expected, account.ID)
			}
		})
	}
}

func TestUpdateCreditLimit(t *testing.T) {
	updater := NewAccountCreditLimitUpdater(NewConnection("127.0.0.1", 5432, "docker", "docker", "routine"))
	var tests = []struct {
		name     string
		expected error
		given    *models.Account
	}{
		{"invalid id", sql.ErrNoRows, new(models.Account)},
		{"success", nil, &models.Account{ID: 1, AvailableCreditLimit: 100}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := updater.UpdateCreditLimit(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
