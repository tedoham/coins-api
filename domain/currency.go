package domain

type CurrencyType struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
