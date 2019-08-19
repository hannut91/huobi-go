package types

type Symbol struct {
	BaseCurrency    string  `json:"base-currency"`
	QuoteCurrency   string  `json:"quote-currency"`
	PricePrecision  int     `json:"price-precision"`
	AmountPrecision int     `json:"amount-precision"`
	SymbolPartition string  `json:"symbol-partition"`
	Symbol          string  `json:"symbol"`
	State           string  `json:"state"`
	ValuePrecision  int     `json:"value-precision"`
	MinOrderAmt     float64 `json:"min-order-amt"`
	MaxOrderAmt     float64 `json:"max-order-amt"`
	MinOrderCalue   float64 `json:"min-order-calue"`
}
