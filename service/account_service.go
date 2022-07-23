package service

import (
	"hexagonal-architecture/errs"
	"hexagonal-architecture/logs"
	"hexagonal-architecture/repository"
	"strings"
	"time"
)

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return accountService{accountRepository: accountRepository}
}

func (service accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// Validation
	if request.Amount < 500 {
		return nil, errs.NewValidationError("amount at least 500")
	}
	if strings.ToLower(request.AccountType) != "saving" || strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAccount, err := service.accountRepository.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAccount.AccountID,
		OpeningDate: newAccount.OpeningDate,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      newAccount.Status,
	}

	return &response, nil
}

func (service accountService) GetAccount(customerID int) ([]AccountResponse, error) {
	accounts, err := service.accountRepository.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}
