package repository

import "errors"

//demonstration purpose, to show new plung
type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerId: 3, Name: "Robbie", City: "Bkk", ZipCode: "333", Status: 1, DateOfBirth: "12-12-2012"},
		{CustomerId: 4, Name: "Bryant", City: "CA", ZipCode: "311", Status: 1, DateOfBirth: "12-12-2012"},

	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerId == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}