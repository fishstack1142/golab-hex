package repository

import "github.com/jmoiron/sqlx"

//if CustomerRepositoryDB == public
// if customerRepositoryDB == private
type customerRepositoryDB struct {
	db *sqlx.DB
}

//for new instance
//constructor, that is
func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

//adapter
func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
//adapter
func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	
	return &customer, nil
}