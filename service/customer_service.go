package service

import (
	"database/sql"
	"errors"
	"hexagonal-architecture/repository"
	"log"
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
		log.Println(err)
		return nil, err
	}

	customerReponses := []CustomerResponse{}
	for _, customer := range customers {
		customerReponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerReponses = append(customerReponses, customerReponse)
	}

	return customerReponses, nil
}

// Get customer by ID
func (service *customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := service.customerRepository.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
