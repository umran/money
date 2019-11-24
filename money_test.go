package money

import (
	"testing"
)

func TestNew(t *testing.T) {
	amount := int64(20000)
	m := New(amount, USD)

	if m.Amount() != amount {
		t.Error("unexpected amount")
	}
}

func TestCurrency(t *testing.T) {
	m := New(2000, USD)
	if m.Currency().Code() != USD.Code() {
		t.Error("unexpected currency")
		t.Log(m.Currency().Code())
	}
}

func TestAdd(t *testing.T) {
	m1 := New(2000, USD)
	m2 := New(1000, USD)

	result, _ := m1.Add(m2)

	if result.Amount() != 3000 {
		t.Error("unexpected result")
		t.Log(result.Amount())
	}
}

func TestSub(t *testing.T) {
	m1 := New(2000, USD)
	m2 := New(1000, USD)

	result, _ := m1.Sub(m2)

	if result.Amount() != 1000 {
		t.Error("unexpected result")
		t.Log(result.Amount())
	}
}

func TestMul(t *testing.T) {
	m := New(100, USD)

	result := m.Mul(0.3333333)

	if result.Amount() != 33 {
		t.Error("unexpected result")
		t.Log(result.Amount())
	}
}

func TestInvalidAdd(t *testing.T) {
	m1 := New(20000, USD)
	m2 := New(10000, CAD)

	_, err := m1.Add(m2)
	if err != ErrCurrencyCodeMismatch {
		t.Error("unexpected error")
		t.Log(err)
	}
}

func TestInvalidSub(t *testing.T) {
	USDA := &Currency{
		code:     "USD",
		exponent: 2,
	}

	USDB := &Currency{
		code:     "USD",
		exponent: 3,
	}

	m1 := New(20000, USDA)
	m2 := New(10000, USDB)

	_, err := m1.Sub(m2)
	if err != ErrCurrencyExponentMismatch {
		t.Error("unexpected error")
		t.Log(err)
	}
}
