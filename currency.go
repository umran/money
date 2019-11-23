package money

// Currency represents a currency and its exponent
type Currency struct {
	code     string
	exponent int
}

// Exponent returns the exponent
func (c *Currency) Exponent() int {
	return c.exponent
}

// Code returns the currency code
func (c *Currency) Code() string {
	return c.code
}
