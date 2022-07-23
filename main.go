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
	customerRepository := repository.NewCustomerRepositoryDB(db)
	accountRepository := repository.NewAccountRepositoryDB(db)

	// Service
	customerService := service.NewCustomerService(customerRepository)
	accountService := service.NewAccountService(accountRepository)

	// Handler
	customerHandler := handler.NewCustomerHandler(&customerService)
	accountHandler := handler.NewAccountHandler(accountService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/accounts/{customerID:[0-9]+}", accountHandler.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)

	// Run server
	port := fmt.Sprintf(":%v", viper.GetInt("app.port"))
	logs.Info("Server running at http://localhost" + viper.GetString("app.port"))
	http.ListenAndServe(port, router)
}
