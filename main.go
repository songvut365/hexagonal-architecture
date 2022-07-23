package main

import (
	"fmt"
	"hexagonal-architecture/config"
	"hexagonal-architecture/handler"
	"hexagonal-architecture/logs"
	"hexagonal-architecture/repository"
	"hexagonal-architecture/service"
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
	_ = customerRepositoryMock

	// Service
	customerService := service.NewCustomerService(customerRepository)

	// Handler
	customerHandler := handler.NewCustomerHandler(&customerService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	// Run server
	port := fmt.Sprintf(":%v", viper.GetInt("app.port"))
	logs.Info("Server running at http://localhost" + viper.GetString("app.port"))
	http.ListenAndServe(port, router)
}
