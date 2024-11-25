package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// related to the stringer interface in the fmt package.
type Stringer interface {
	String() string
}

// stringer method when you call the %s in the fmt library.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("insufficient balance")
	}
	w.balance -= amount
	return nil
}
