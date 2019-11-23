package money

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Print(ErrCurrencyCodeMismatch)
}
