package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

var app *App

func TestMain(m *testing.M) {
	app = NewApp(
		new(usecasesMock),
		new(usecasesMock),
		new(usecasesMock))

	code := m.Run()

	os.Exit(code)
}

type usecasesMock struct{}

// CreateAccount register an account
func (u *usecasesMock) CreateAccount(acc *models.Account) error {
	acc.ID++
	return nil
}

// GetAccount search by account with supplied identifier
func (u *usecasesMock) GetAccount(accID int) (*models.Account, error) {
	return &models.Account{ID: 1, DocumentNumber: "99999999900"}, nil
}

// CreateTransaction register a transaction
func (u *usecasesMock) CreateTransaction(transaction *models.Transaction) error {
	transaction.ID++
	return nil
}

func TestCreateAccount(t *testing.T) {
	var tests = []struct {
		name             string
		expected         int
		givenContentType string
		givenAccount     *models.Account
	}{
		{"empty content type", http.StatusUnsupportedMediaType, "", nil},
		{"empty account", http.StatusBadRequest, "application/json", nil},
		{"success", http.StatusOK, "application/json", &models.Account{DocumentNumber: "99999999900"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			err := json.NewEncoder(buf).Encode(tt.givenAccount)
			if err != nil {
				t.Fatalf("fail to encode json: %+v", err)
			}

			// create response recorder
			rr := httptest.NewRecorder()
			// creating request
			req := httptest.NewRequest("POST", "/accounts", buf)
			req.Header.Set("Content-Type", tt.givenContentType)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			app.Router.ServeHTTP(rr, req)
			if actual := rr.Code; actual != tt.expected {
				t.Errorf("(%s, %+v): expected %d, actual %d", tt.givenContentType, tt.givenAccount, tt.expected, actual)
			}

		})
	}
}

func TestGetAccount(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    string
	}{
		{"invalid id", http.StatusNotFound, ""},
		{"success", http.StatusOK, "1"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// create response recorder
			rr := httptest.NewRecorder()
			// creating request
			req := httptest.NewRequest("GET", fmt.Sprintf("/accounts/%s", tt.given), nil)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			app.Router.ServeHTTP(rr, req)
			if actual := rr.Code; actual != tt.expected {
				t.Errorf("(%s): expected %d, actual %d", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestCreateTransaction(t *testing.T) {
	var tests = []struct {
		name             string
		expected         int
		givenContentType string
		givenTransaction *models.Transaction
	}{
		{"empty content type", http.StatusUnsupportedMediaType, "", nil},
		{"empty transaction", http.StatusBadRequest, "application/json", nil},
		{"success", http.StatusOK, "application/json", &models.Transaction{
			Account: models.Account{
				ID: 1,
			},
			Operation: 0,
			Amount:    0.0,
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			err := json.NewEncoder(buf).Encode(tt.givenTransaction)
			if err != nil {
				t.Fatalf("fail to encode json: %+v", err)
			}

			// create response recorder
			rr := httptest.NewRecorder()
			// creating request
			req := httptest.NewRequest("POST", "/transactions", buf)
			req.Header.Set("Content-Type", tt.givenContentType)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			app.Router.ServeHTTP(rr, req)
			if actual := rr.Code; actual != tt.expected {
				t.Errorf("(%s, %+v): expected %d, actual %d", tt.givenContentType, tt.givenTransaction, tt.expected, actual)
			}

		})
	}
}
