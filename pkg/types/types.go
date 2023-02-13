package types

type Currency string

const (
	TJS Currency = "TJK"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    int
	MinBalance int
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}
type Payment struct {
	ID     int
	Amount int
}
