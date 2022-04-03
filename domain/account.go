package domain

type Account struct {
	ID            int     `db:"id"`
	Name          string  `db:"name"`
	AccountNumber string  `db:"name"`
	Balance       float64 `db:"balance"`
}
