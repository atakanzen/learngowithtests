package pointers_errors_test

import (
	pe "learngowithtests/pointers_errors"
	"math/rand"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := pe.Wallet{}
		wallet.Deposit(pe.Bitcoin(22))

		assertBalance(t, wallet, pe.Bitcoin(22))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := pe.Wallet{}

		// Since it's a blackbox test strategy, deposit is needed here.
		wallet.Deposit(pe.Bitcoin(10))
		err := wallet.Withdraw(pe.Bitcoin(2))

		assertNoErr(t, err)
		assertBalance(t, wallet, pe.Bitcoin(8))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := pe.Wallet{}

		wallet.Deposit(pe.Bitcoin(17))
		err := wallet.Withdraw(pe.Bitcoin(100))

		assertErr(t, err, pe.ErrInsufficientFunds)
		assertBalance(t, wallet, pe.Bitcoin(17))
	})
}

func ExampleWallet_Deposit() {
	wallet := pe.Wallet{}
	wallet.Deposit(pe.Bitcoin(22))
	//Output:

}

func ExampleWallet_Withdraw() {
	wallet := pe.Wallet{}
	wallet.Withdraw(pe.Bitcoin(10))
	//Output:

}

func ExampleWallet_Balance() {
	wallet := pe.Wallet{}
	wallet.Balance()
	//Output:

}

func BenchmarkWallet_Deposit(b *testing.B) {
	wallet := pe.Wallet{}
	for i := 0; i < b.N; i++ {
		wallet.Deposit(pe.Bitcoin(rand.Intn(b.N)))
	}
}

func assertBalance(t testing.TB, wallet pe.Wallet, want pe.Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoErr(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("err must be nil")
	}
}

func assertErr(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("err must not be nil")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got.Error(), want)
	}
}
