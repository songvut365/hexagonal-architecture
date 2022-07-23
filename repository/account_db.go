package repository

import (
	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (repository accountRepositoryDB) Create(account Account) (*Account, error) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := repository.db.Exec(
		query,
		account.CustomerID,
		account.OpeningDate,
		account.AccountType,
		account.Amount,
		account.Status,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	account.AccountID = int(id)

	return &account, nil
}

func (repository accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE customer_id = ?"
	accounts := []Account{}

	err := repository.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
