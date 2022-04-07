package domain

import (
	"github.com/tedoham/coins-api/dto"
	"github.com/tedoham/coins-api/errs"
)

type Account struct {
	AccountId int     `db:"id"`
	Name      string  `db:"name"`
	Balance   float64 `db:"balance"`
}

type AccountList struct {
	AccountId string `db:"id"`
	Name      string `db:"name"`
	Balance   string `db:"balance"`
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Balance < amount {
		return false
	}
	return true
}

func (t AccountList) ToDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountId: t.AccountId,
		Name:      t.Name,
		Balance:   t.Balance,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockAccounRepository.go -package=mock github.com/tedoham/coins-api/domain AccountRepository
type AccountRepository interface {
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
	FindAllPayments() ([]Transaction, *errs.AppError)
	FindAllAccounts() ([]AccountList, *errs.AppError)
}
