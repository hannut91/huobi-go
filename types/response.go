package types

type Response struct {
	Status       string `json:"status"`
	Channel      string `json:"ch"`
	Timestamp    int64  `json:"ts"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
}

type BalanceResponse struct {
	ID    int            `json:"id"`
	Type  string         `json:"type"`
	State string         `json:"state"`
	List  []*BalanceData `json:"list"`
}

type BalanceData struct {
	Currency string `json:"currency"`
	Type     string `json:"type"`
	Balance  string `json:"balance"`
}
