package domain

import (
	"github.com/tedoham/coins-api/dto"
)

type Transaction struct {
	TransactionId   string  `db:"id"`
	FromAccount     string  `db:"from_account"`
	ToAccount       string  `db:"to_account"`
	TransactionType string  `db:"transaction_type"`
	CurrencyType    string  `db:"currency_type"`
	Amount          float64 `db:"amount"`
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		FromAccount:     t.FromAccount,
		ToAccount:       t.ToAccount,
		TransactionType: t.TransactionType,
		CurrencyType:    t.CurrencyType,
		Amount:          t.Amount,
	}
}

const WITHDRAWAL = "outgoing"

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
