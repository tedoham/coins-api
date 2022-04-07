package service

import (
	"github.com/tedoham/coins-api/domain"
	"github.com/tedoham/coins-api/dto"
	"github.com/tedoham/coins-api/errs"
)

//go:generate mockgen -destination=../mocks/service/mockAccountService.go -package=mock github.com/tedoham/coins-api/service AccountService
type AccountService interface {
	ListAllAccount() ([]dto.AccountResponse, *errs.AppError)
	ListAllPayments() ([]dto.TransactionResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) ListAllAccount() ([]dto.AccountResponse, *errs.AppError) {
	accounts, err := s.repo.FindAllAccounts()

	if err != nil {
		return nil, err
	}
	response := make([]dto.AccountResponse, 0)
	for _, c := range accounts {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultAccountService) ListAllPayments() ([]dto.TransactionResponse, *errs.AppError) {
	payments, err := s.repo.FindAllPayments()
	if err != nil {
		return nil, err
	}
	response := make([]dto.TransactionResponse, 0)
	for _, c := range payments {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account
	account, err := s.repo.FindBy(req.FromAccount)
	if err != nil {
		return nil, err
	}
	if !account.CanWithdraw(req.Amount) {
		return nil, errs.NewValidationError("Insufficient balance in the account")
	}

	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		FromAccount:     req.FromAccount,
		ToAccount:       req.ToAccount,
		TransactionType: "outgoing",
		CurrencyType:    "USD",
		Amount:          req.Amount,
	}

	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
