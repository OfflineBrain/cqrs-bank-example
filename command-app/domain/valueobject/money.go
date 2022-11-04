package valueobject

type Price struct {
	Currency Currency
	Value    CurrencyValue
}

type Currency string

type CurrencyValue float32
