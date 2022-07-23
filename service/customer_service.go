package service

import (
	"database/sql"
	"errors"
	"hexagonal-architecture/logs"
	"hexagonal-architecture/repository"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

// constructor
func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{customerRepository: customerRepository}
}

// Get customers
func (service *customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := service.customerRepository.GetAll()
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses, nil
}

// Get customer by ID
func (service *customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := service.customerRepository.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		logs.Error(err)
		return nil, err
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
