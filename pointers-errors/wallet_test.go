package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		checkBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		err := wallet.Withdraw(Bitcoin(10))
		checkNoError(t, err)
		checkBalance(t, wallet, 10)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(40))
		checkError(t, err, errInsufficient)
		checkBalance(t, wallet, 20)

	})
}

func checkBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s\nwant %s", got, want)
	}
}

func checkError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func checkNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't get one")
	}
}
