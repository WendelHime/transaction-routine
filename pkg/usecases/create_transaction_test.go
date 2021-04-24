package usecases

import (
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
	"github.com/WendelHime/transaction-routine/pkg/domain/services"
)

type tranRepo struct{}

// Create register a transaction on persistence service
func (r *tranRepo) Create(transaction *models.Transaction) error {
	return nil
}

type accRepo struct{}

// Create register an account on persistence service
func (u *accRepo) Create(account *models.Account) error {
	return nil
}

// Get searches by account using the supplied identifier and returns
// a account reference from persistence service
func (u *accRepo) Get(accountID int) (*models.Account, error) {
	return &models.Account{
		ID:                   1,
		AvailableCreditLimit: 100,
	}, nil
}

// UpdateCreditLimit update credit limit from account on persistence service
func (u *accRepo) UpdateCreditLimit(account *models.Account) error {
	return nil
}

func TestCreateTransaction(t *testing.T) {
	uc := NewUCTransactionCreator(services.NewTransactionService(new(tranRepo)), services.NewAccountService(new(accRepo), new(accRepo), new(accRepo)))
	var tests = []struct {
		name     string
		expected error
		given    *models.Transaction
	}{
		{"account id invalid", models.ErrMissingRequiredField, &models.Transaction{
			Account: models.Account{
				ID:                   0,
				AvailableCreditLimit: 0.0,
			},
			Operation: 0,
			Amount:    0.0,
		}},
		{"amount above credit limit", models.ErrAmountBeyondCreditLimit, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: models.OperationTypeCompraVista,
			Amount:    -110,
		}},
		{"credit equal to amount", nil, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: models.OperationTypeCompraVista,
			Amount:    -100,
		}},
		{"amount lesser than credit limit", nil, &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: models.OperationTypeCompraVista,
			Amount:    -20,
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := uc.CreateTransaction(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
