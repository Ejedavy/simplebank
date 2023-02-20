package db

import (
	"context"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"simple_bank/db/utils"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	param := CreateAccountParams{
		Owner:    utils.RandomName(6),
		Balance:  utils.RandomNumber(0, 1000),
		Currency: utils.RandomCurrency(),
	}

	acct, err := testQueries.CreateAccount(context.Background(), param)
	require.NotEmpty(t, acct)
	require.NoError(t, err)
	require.Equal(t, param.Owner, acct.Owner)
	require.Equal(t, param.Balance, acct.Balance)
	require.Equal(t, param.Currency, acct.Currency)
	require.NotZero(t, acct.CreatedAt)
	require.NotZero(t, acct.ID)

	return acct
}

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	account := createRandomAccount(t)
	gottenAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NotEmpty(t, gottenAccount)
	require.NoError(t, err)

	require.Equal(t, account.ID, gottenAccount.ID)
	require.Equal(t, account.Balance, gottenAccount.Balance)
	require.Equal(t, account.Owner, gottenAccount.Owner)
	require.Equal(t, account.Currency, gottenAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, gottenAccount.CreatedAt, time.Second)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	gottenAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Empty(t, gottenAccount)
	require.Error(t, err)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	oldBalance := account.Balance
	credit := utils.RandomNumber(100, 1000)
	params := UpdateAccountParams{
		ID:      account.ID,
		Balance: credit,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
	require.Equal(t, oldBalance+credit, account2.Balance)
	require.NotEqual(t, oldBalance, account2.Balance)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	params := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
