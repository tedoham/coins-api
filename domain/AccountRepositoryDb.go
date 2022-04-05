package domain

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/tedoham/coins-api/errs"
	"github.com/tedoham/coins-api/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

/**
 * transaction = make an entry in the transaction table + update the balance in the accounts table
 */
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction for outgoing
	var id int64
	result := tx.QueryRow(`INSERT INTO transactions (from_account, to_account, transaction_type, currency_type, amount)
		values ($1, $2, $3, $4, $5) RETURNING id`, t.FromAccount, t.ToAccount, t.TransactionType, t.CurrencyType, t.Amount).Scan(&id)
	_, err = tx.Exec(`UPDATE accounts SET balance = balance - $1 where id = $2`, t.Amount, t.FromAccount)

	// inserting bank account transaction for incoming
	t.TransactionType = "incoming"
	_, err = tx.Exec(`INSERT INTO transactions (from_account, to_account, transaction_type, currency_type, amount)
		values ($1, $2, $3, $4, $5)`, t.FromAccount, t.ToAccount, t.TransactionType, t.CurrencyType, t.Amount)
	_, err = tx.Exec(`UPDATE accounts SET balance = balance + $1 where id = $2`, t.Amount, t.ToAccount)

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	if result != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.FromAccount)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(id, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Balance
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := fmt.Sprintf(`SELECT id, name, balance from accounts where id = $1`)

	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func (d AccountRepositoryDb) FindAllPayments() ([]Transaction, *errs.AppError) {
	var err error
	payments := make([]Transaction, 0)

	findAllSql := fmt.Sprintf(`
		SELECT 
		T.id, 
		A.name as from_account, 
		B.name as to_account, 
		transaction_type, 
		currency_type, 
		amount 
		from transactions AS T
		JOIN accounts AS A ON A.id = T.from_account
		JOIN accounts AS B ON B.id = T.to_account
	`)

	err = d.client.Select(&payments, findAllSql)

	if err != nil {
		logger.Error("Error while querying account table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return payments, nil
}

func (d AccountRepositoryDb) FindAllAccounts() ([]AccountList, *errs.AppError) {
	var err error
	accounts := make([]AccountList, 0)

	findAllSql := "SELECT id, name, balance from accounts"

	err = d.client.Select(&accounts, findAllSql)

	if err != nil {
		logger.Error("Error while querying account table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return accounts, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
