package postgres

import (
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

func TestCreateTransaction(t *testing.T) {
	creator := NewTransactionCreator(NewConnection("127.0.0.1", 5432, "docker", "docker", "routine"))
	var tests = []struct {
		name     string
		expected error
		given    *models.Transaction
	}{
		{"success", nil, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: 4,
			Amount:    0.0,
		}},
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
