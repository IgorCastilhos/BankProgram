package utils

// Constants for all supported currencies
const (
	BRL = "BRL"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case BRL:
		return true
	}
	return false
}
