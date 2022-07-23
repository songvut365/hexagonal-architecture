package service

import (
	"database/sql"
	"hexagonal-architecture/errs"
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
		return nil, errs.NewUnexpectedError()
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
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
