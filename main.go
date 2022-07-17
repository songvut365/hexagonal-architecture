package main

import (
	"fmt"
	"hexagonal-architecture/handler"
	"hexagonal-architecture/repository"
	"hexagonal-architecture/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	// Database
	db, err := sqlx.Open("mysql", "root:1234@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	// Repository
	customerRepository := repository.NewCustomerRepositoryDB(db)

	// Change repository to mock
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	// customerService := service.NewCustomerService(customerRepositoryMock)

	// Service
	customerService := service.NewCustomerService(customerRepository)

	// Handler
	customerHandler := handler.NewCustomerHandler(&customerService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}

func TestRepository(db *sqlx.DB) {
	customerRepository := repository.NewCustomerRepositoryDB(db)

	customersFromRepo, err := customerRepository.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(customersFromRepo)

	customerFromRepo, err := customerRepository.GetById(1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(customerFromRepo)
}

func TestService(customerService service.CustomerService) {
	customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)

	customer, err := customerService.GetCustomer(1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
