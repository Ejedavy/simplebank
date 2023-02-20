package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simple_bank/db/utils"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	params := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomNumber(0, 400),
	}

	entry, err := testQueries.CreateEntry(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	require.Equal(t, params.AccountID, entry.AccountID)
	require.Equal(t, params.Amount, entry.Amount)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestQueries_ListAccountEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	params := ListAccountEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListAccountEntries(context.Background(), params)
	require.Len(t, entries, 5)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestQueries_GetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	gottenEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gottenEntry)
	require.NotZero(t, gottenEntry.ID)
	require.WithinDuration(t, entry.CreatedAt, gottenEntry.CreatedAt, time.Second)
	require.Equal(t, gottenEntry.ID, entry.ID)
	require.Equal(t, gottenEntry.Amount, entry.Amount)
	require.Equal(t, gottenEntry.AccountID, entry.AccountID)
}

func TestQueries_ListBankEntries(t *testing.T) {
	accounts := make([]Account, 0, 10)
	for i := 0; i < 10; i++ {
		accounts = append(accounts, createRandomAccount(t))
	}

	for _, account := range accounts {
		createRandomEntry(t, account)
	}

	params := ListBankEntriesParams{
		Limit:  10,
		Offset: 10,
	}

	entries, err := testQueries.ListBankEntries(context.Background(), params)
	require.Len(t, entries, 10)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
