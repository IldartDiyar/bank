package db

import (
	"context"
	"testing"
	"time"

	"github.com/IldartDiyar/bank/util"
	"github.com/stretchr/testify/require"
)

func createRamdomAccount(t *testing.T) Account {
	arg := CreateAccountParams{util.RandomOwner(), util.RamdomMoney(), util.RandomCurrency()}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRamdomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRamdomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRamdomAccount(t)

	arg := UpdateAccountParams{
		account1.ID,
		util.RamdomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRamdomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var accounts []Account
	for i := 0; i < 5; i++ {
		accounts = append(accounts, createRamdomAccount(t))
	}
	arg := ListAccountsParams{Limit: 5, Offset: 5}

	accounts1, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	require.Len(t, accounts, 5)

	for i, acc := range accounts1 {
		require.NotEmpty(t, acc)
		require.Equal(t, acc, accounts[i])
	}
}

func TestAddAccountBalance(t *testing.T) {
	account1 := createRamdomAccount(t)

	arg := AddAccountBalanceParams{Amount: util.RamdomMoney(), ID: account1.ID}

	account2, err := testQueries.AddAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance+arg.Amount, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}
