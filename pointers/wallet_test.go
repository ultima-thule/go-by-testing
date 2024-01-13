package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(5))

		asserNoError(t, err)
		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(Bitcoin(25))

		asserError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func asserError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func asserNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("didn't want an error but get one")
	}
}
