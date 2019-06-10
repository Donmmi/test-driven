package main

import (
	"fmt"
	"errors"
)

var InsufficientError = errors.New("insufficient withdraw")

type BitCoin int

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BitCoin", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(d BitCoin) {
	w.balance = w.balance + d
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(d BitCoin) error {
	if w.balance < d {
		return InsufficientError
	}
	w.balance -= d
	return nil
}

func main() {

}
