// Package Pointers & Errors.
package pointers_errors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds string = "cannot withdraw, insufficient funds"

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit the specified amount to the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Return current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw the specified amount from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New(ErrInsufficientFunds)
	}

	w.balance -= amount
	return nil
}
