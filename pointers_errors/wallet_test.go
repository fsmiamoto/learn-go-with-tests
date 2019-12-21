package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("expected no error but got one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(10)
		assertBalance(t, w, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{balance: 25}
		err := w.Withdraw(5)
		assertBalance(t, w, 20)
		assertNoError(t, err)
	})

	t.Run("Withdraw without sufficient funds", func(t *testing.T) {
		prevBalance := Bitcoin(25)

		w := Wallet{balance: prevBalance}
		err := w.Withdraw(35)

		assertBalance(t, w, prevBalance)
		assertError(t, err, ErrInsufficientFunds)

	})

}
