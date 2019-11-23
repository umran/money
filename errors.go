package money

import "fmt"

type genericError struct {
	code    string
	message string
}

// Error implements the error interface
func (ge *genericError) Error() string {
	return fmt.Sprintf("%s: %s", ge.code, ge.message)
}

var (
	// ErrCurrencyCodeMismatch ...
	ErrCurrencyCodeMismatch = &genericError{
		code:    "1001",
		message: "currency code mismatch",
	}

	// ErrCurrencyExponentMismatch ...
	ErrCurrencyExponentMismatch = &genericError{
		code:    "1002",
		message: "currency exponent mismatch",
	}
)
