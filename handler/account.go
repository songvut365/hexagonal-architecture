package handler

import (
	"encoding/json"
	"hexagonal-architecture/errs"
	"hexagonal-architecture/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) accountHandler {
	return accountHandler{
		accountService: accountService,
	}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	response, err := h.accountService.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	account, err := h.accountService.GetAccount(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application")
	json.NewEncoder(w).Encode(account)
}
