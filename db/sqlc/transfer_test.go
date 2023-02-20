package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simple_bank/db/utils"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T, from Account, to Account) Transfer {
	params := TransferTXParam{
		SenderID:   from.ID,
		ReceiverID: to.ID,
		Amount:     utils.RandomNumber(0, 500),
	}
	result, err := testStore.TransferTX(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.Transfer.Amount, params.Amount)
	require.Equal(t, result.Transfer.FromAccountID, params.SenderID)
	require.Equal(t, result.Transfer.ToAccountID, params.ReceiverID)
	require.NotZero(t, result.Transfer.ID)
	require.NotZero(t, result.Transfer.CreatedAt)

	return result.Transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_GetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, account1, account2)
	gottenTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gottenTransfer)

	require.Equal(t, transfer.FromAccountID, gottenTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, gottenTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, gottenTransfer.Amount)

	require.Equal(t, transfer.ID, gottenTransfer.ID)
	require.WithinDuration(t, transfer.CreatedAt, gottenTransfer.CreatedAt, time.Second)

}
func TestQueries_ListIncomingTransfers(t *testing.T) {
	receiver := createRandomAccount(t)

	var senders []Account
	for i := 0; i < 10; i++ {
		senders = append(senders, createRandomAccount(t))
	}

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, senders[i], receiver)
	}
	params := ListIncomingTransfersParams{
		ToAccountID: receiver.ID,
		Limit:       5,
		Offset:      5,
	}
	transfers, err := testQueries.ListIncomingTransfers(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, transfer.ToAccountID, receiver.ID)
		require.Equal(t, transfer.ToAccountID, params.ToAccountID)
	}

}
func TestQueries_ListTransfersBetween(t *testing.T) {

	var senders []Account
	for i := 0; i < 5; i++ {
		senders = append(senders, createRandomAccount(t))
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == i {
				continue
			}
			createRandomTransfer(t, senders[i], senders[j])
		}
	}

	params := ListTransfersBetweenParams{
		Account1: senders[2].ID,
		Account2: senders[4].ID,
		Limit:    5,
		Offset:   0,
	}

	transfers, err := testQueries.ListTransfersBetween(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 2)

	for _, transfer := range transfers {
		require.True(t, transfer.FromAccountID == senders[2].ID || transfer.ToAccountID == senders[2].ID)
		require.True(t, transfer.FromAccountID == senders[4].ID || transfer.ToAccountID == senders[4].ID)
	}

}

func TestQueries_ListOutgoingTransfers(t *testing.T) {
	sender := createRandomAccount(t)

	var receivers []Account
	for i := 0; i < 10; i++ {
		receivers = append(receivers, createRandomAccount(t))
	}

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, sender, receivers[i])
	}
	params := ListOutgoingTransfersParams{
		FromAccountID: sender.ID,
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListOutgoingTransfers(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, transfer.FromAccountID, sender.ID)
		require.Equal(t, transfer.FromAccountID, params.FromAccountID)
	}
}
