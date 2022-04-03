package domain

type Payment struct {
	ID           int     `db:"id"`
	FromAccount  int     `db:"from_account"`
	ToAccount    int     `db:"to_account"`
	CurrencyType int     `db:"to_account"`
	Direction    string  `db:"direction"`
	Amount       float64 `db:"balance"`
}
