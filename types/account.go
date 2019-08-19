package types

type AggregateBalance struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance,string"`
}

type Account struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	State  string `json:"state"`
	UserID int    `json:"user-id"`
}

type Balance struct {
	Locked    string
	Available string
}
