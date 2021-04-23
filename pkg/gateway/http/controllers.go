// Package http contain all HTTP operations
package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WendelHime/transaction-routine/pkg/domain/models"
	"github.com/WendelHime/transaction-routine/pkg/usecases"
	"github.com/gorilla/mux"
)

type controller interface {
	// Accounts create new accounts if receive a POST request,
	// search by account with supplied ID if receive a GET request.
	Accounts(w http.ResponseWriter, r *http.Request)
	// Transactions create new transactions if receive a POST request
	Transactions(w http.ResponseWriter, r *http.Request)
}

// App the http app to be executed
type App struct {
	Router             *mux.Router
	accountCreator     usecases.UCAccountCreator
	accountGetter      usecases.UCAccountGetter
	transactionCreator usecases.UCTransactionCreator
}

// NewApp build a new app and initialize routes
func NewApp(accountCreator usecases.UCAccountCreator, accountGetter usecases.UCAccountGetter, transactionCreator usecases.UCTransactionCreator) *App {
	app := App{accountCreator: accountCreator, accountGetter: accountGetter, transactionCreator: transactionCreator}
	app.Router = mux.NewRouter()
	app.initializeRoutes()
	return &app
}

func (h *App) getAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, err := strconv.Atoi(params["accountID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}

	acc, err := h.accountGetter.GetAccount(accountID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error: %v", err)
		return
	}
	account := &Account{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error: %v", err)
		return
	}

}

func (h *App) createAccount(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") == "" || r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		fmt.Fprint(w, "unsupported media type")
		return
	}

	var acc *Account
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		log.Printf("fail parsing account json: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}

	if acc == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}

	account := &models.Account{DocumentNumber: acc.DocumentNumber}
	err = h.accountCreator.CreateAccount(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error: %v", err)
		return
	}
	acc.ID = account.ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(acc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error: %v", err)
		return
	}

}

func (h *App) createTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "" || r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		fmt.Fprint(w, "unsupported media type")
		return
	}
	var trans *Transaction
	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Printf("fail parsing transaction json: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}

	if trans == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}
	transaction := &models.Transaction{
		Account: models.Account{
			ID: trans.AccountID,
		},
		Operation: trans.Operation,
		Amount:    trans.Amount,
	}
	err = h.transactionCreator.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error: %v", err)
		return
	}

	trans.ID = transaction.ID
	trans.EventDate = transaction.EventDate
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(trans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error: %v", err)
		return
	}
}

func (h *App) initializeRoutes() {
	h.Router.HandleFunc("/accounts", h.createAccount).Methods("POST")
	h.Router.HandleFunc("/accounts/{accountID:[0-9]+}", h.getAccount).Methods("GET")
	h.Router.HandleFunc("/transactions", h.createTransaction).Methods("POST")
}

// Run http server
func (h *App) Run(addr string) {
	log.Printf("listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, h.Router))
}
