package postgres

import (
	"database/sql"
	"log"

	"github.com/WendelHime/transaction-routine/pkg/domain/interfaces"
	"github.com/WendelHime/transaction-routine/pkg/domain/models"
)

type accountRepository struct {
	conn *Connection
}

// NewAccountCreator builds a new account creator implementation
func NewAccountCreator(conn *Connection) interfaces.AccountCreator {
	return &accountRepository{conn: conn}
}

// NewAccountGetter builds a new account getter implementation
func NewAccountGetter(conn *Connection) interfaces.AccountGetter {
	return &accountRepository{conn: conn}
}

// NewAccountCreditLimitUpdater builds a new implementation for cred limit update
func NewAccountCreditLimitUpdater(conn *Connection) interfaces.CreditLimitUpdater {
	return &accountRepository{conn: conn}
}

// Create register an account on persistence service
func (r *accountRepository) Create(account *models.Account) error {
	db, err := r.conn.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(
		`INSERT INTO accounts(document_number, available_credit_limit)
		VALUES ($1, $2) RETURNING id`,
		account.DocumentNumber, account.AvailableCreditLimit).Scan(&account.ID)
	if err != nil {
		return err
	}

	return nil
}

// Get searches by account using the supplied identifier and returns
// a account reference from persistence service
func (r *accountRepository) Get(accountID int) (*models.Account, error) {
	db, err := r.conn.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := db.QueryRow(
		`SELECT
			id,
			document_number,
			available_credit_limit
		FROM accounts
		WHERE id=$1`, accountID)

	account := new(models.Account)
	switch err := row.Scan(
		&account.ID,
		&account.DocumentNumber,
		&account.AvailableCreditLimit); err {
	case sql.ErrNoRows:
		log.Println("no rows")
		return new(models.Account), err
	case nil:
		return account, nil
	default:
		log.Println(err)
		return new(models.Account), err
	}
}

// UpdateCreditLimit update credit limit from account on persistence service
func (r *accountRepository) UpdateCreditLimit(account *models.Account) error {
	db, err := r.conn.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(
		`UPDATE accounts SET available_credit_limit=$1 WHERE id=$2 RETURNING id`,
		account.AvailableCreditLimit, account.ID).Scan(&account.ID)

	return err
}
