package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBitCoin(t, wallet.Balance(), BitCoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{BitCoin(100)}
		err := wallet.Withdraw(BitCoin(10))

		assertNoError(t, err)
		assertBitCoin(t, wallet.Balance(), BitCoin(90))
	})

	t.Run("withdraw with insufficient", func(t *testing.T) {
		wallet := Wallet{BitCoin(10)}
		err := wallet.Withdraw(BitCoin(20))
		assertError(t, err, InsufficientError)
	})
}

func assertBitCoin(t *testing.T, got, want BitCoin) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("should not get an error")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("should get non nil error")
	}

	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}