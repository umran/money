package money

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/cowriepayments/decimal"
)

// Money represents an amount of some currency
type Money struct {
	amount   decimal.Decimal
	currency *Currency
}

// Add performs addition with another Money instance and returns a new Money instance
func (m1 *Money) Add(m2 *Money) (*Money, error) {
	// make sure we're adding the same currency
	if err := assertIdenticalCurrency(m1, m2); err != nil {
		return nil, err
	}

	return &Money{
		amount:   m1.amount.Add(m2.amount),
		currency: m1.currency,
	}, nil
}

// Sub performs subtraction with another Money instance and returns a new Money instance
func (m1 *Money) Sub(m2 *Money) (*Money, error) {
	// make sure we're subtracting the same currency
	if err := assertIdenticalCurrency(m1, m2); err != nil {
		return nil, err
	}

	return &Money{
		amount:   m1.amount.Sub(m2.amount),
		currency: m1.currency,
	}, nil
}

// Mul performs multipliciation with a decimal (supplied as a float64) and returns a new Money instance
func (m1 *Money) Mul(factor float64) *Money {
	return &Money{
		amount:   m1.amount.Mul(decimal.NewFromFloat(factor)),
		currency: m1.currency,
	}
}

// Currency returns the currency
func (m1 *Money) Currency() *Currency {
	return m1.currency
}

// Amount returns the amount in minor units (cents) if the exponent is greater than 0
// else returns the amount in major units
func (m1 *Money) Amount() int64 {
	s := m1.amount.StringFixedBank(int32(m1.currency.exponent))
	s = strings.Replace(s, ".", "", 1)

	i, _ := strconv.ParseInt(s, 10, 64)

	return i
}

func assertIdenticalCurrency(m1, m2 *Money) error {
	switch {
	case m1.currency.code != m2.currency.code:
		return ErrCurrencyCodeMismatch
	case m1.currency.exponent != m2.currency.exponent:
		return ErrCurrencyExponentMismatch
	default:
		return nil
	}
}

// New returns a new Money instance from the provided amount (in minor units) and currency
func New(amount int64, currency *Currency) *Money {
	return &Money{
		amount:   decimal.NewFromBigInt(big.NewInt(amount), int32(-1*currency.exponent)),
		currency: currency,
	}
}
