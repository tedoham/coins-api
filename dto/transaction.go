package dto

import (
	"github.com/tedoham/coins-api/errs"
)

const WITHDRAWAL = "outgoing"
const DEPOSIT = "incoming"

type TransactionRequest struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	FromAccount     string  `json:"from_account"`
	ToAccount       string  `json:"to_account"`
	TransactionType string  `db:"transaction_type"`
	CurrencyType    string  `json:"currency_type"`
	Amount          float64 `json:"amount"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}
