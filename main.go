package main

import (
	"fmt"
	"hexagonal-architecture/config"
	"hexagonal-architecture/handler"
	"hexagonal-architecture/repository"
	"hexagonal-architecture/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// Config
	config.InitTimeZone()
	config.InitViperConfig()
	db := config.InitDatabase()

	// Repository
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerRepository := repository.NewCustomerRepositoryDB(db)
	_ = customerRepository

	// Service
	customerService := service.NewCustomerService(customerRepositoryMock)

	// Handler
	customerHandler := handler.NewCustomerHandler(&customerService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	// Run server
	port := fmt.Sprintf(":%v", viper.GetInt("app.port"))
	log.Printf("Server running at http://localhost%v", port)
	http.ListenAndServe(port, router)
}
