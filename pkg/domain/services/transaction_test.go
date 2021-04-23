package services

import (
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

type tranRepo struct{}

// Create register a transaction on persistence service
func (r *tranRepo) Create(transaction *models.Transaction) error {
	return nil
}

func TestCreateTransaction(t *testing.T) {
	serv := NewTransactionService(new(tranRepo))
	var tests = []struct {
		name     string
		expected error
		given    *models.Transaction
	}{
		{"invalid account id", models.ErrMissingRequiredField, new(models.Transaction)},
		{"invalid amount compra", models.ErrInvalidAmount, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: 1,
			Amount:    10,
		}},
		{"invalid amount pagamento", models.ErrInvalidAmount, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: 4,
			Amount:    -10,
		}},
		{"success", nil, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: 4,
			Amount:    10,
		}},
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
