package repository

import "github.com/jmoiron/sqlx"

// Adapter is struct
type customerRepositoryDB struct {
	db *sqlx.DB
}

// constructor
func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

// Get all
func (repository customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"

	err := repository.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// Get by ID
func (repository customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id=?"

	err := repository.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
