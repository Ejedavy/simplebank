package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simple_bank/db/utils"
	"testing"
)

func TestStore_TransferTXWithFixedAmount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	numberOfConcurrentTransactions := 20
	errorChannel := make(chan error)
	resultChannel := make(chan TransferTXResult)
	amount := int64(20)
	existed := make(map[int]bool)

	for i := 0; i < numberOfConcurrentTransactions; i++ {
		go func() {
			result, err := testStore.TransferTX(context.Background(), TransferTXParam{
				SenderID:   account1.ID,
				ReceiverID: account2.ID,
				Amount:     amount,
			})

			errorChannel <- err
			resultChannel <- result
		}()
	}

	for i := 0; i < numberOfConcurrentTransactions; i++ {
		err := <-errorChannel
		result := <-resultChannel
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, result.SenderEntry.AccountID, account1.ID)
		require.Equal(t, result.ReceiverEntry.AccountID, account2.ID)
		require.Equal(t, result.Sender.ID, account1.ID)
		require.Equal(t, result.Receiver.ID, account2.ID)
		require.Equal(t, result.SenderEntry.AccountID, result.Sender.ID)
		require.Equal(t, result.ReceiverEntry.AccountID, result.Receiver.ID)
		require.Equal(t, -result.Amount, result.SenderEntry.Amount)
		require.Equal(t, result.Amount, result.ReceiverEntry.Amount)
		require.Equal(t, -result.Transfer.Amount, result.SenderEntry.Amount)
		require.Equal(t, result.Transfer.Amount, result.ReceiverEntry.Amount)
		require.NotZero(t, result.SenderEntry.ID)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.ReceiverEntry.ID)
		require.Equal(t, result.Sender.Balance+result.Receiver.Balance,
			(result.Sender.Balance+result.Amount)+(result.Receiver.Balance-result.Amount))
		_, err = testStore.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)
		_, err = testStore.GetEntry(context.Background(), result.SenderEntry.ID)
		require.NoError(t, err)
		_, err = testStore.GetEntry(context.Background(), result.ReceiverEntry.ID)
		require.NoError(t, err)

		diff1 := account1.Balance - result.Sender.Balance
		diff2 := result.Receiver.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff2 > 0)

		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= numberOfConcurrentTransactions)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balance
	updatedAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(numberOfConcurrentTransactions)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(numberOfConcurrentTransactions)*amount, updatedAccount2.Balance)

}

func TestStore_TransferTXWithRandomAmounts(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	numberOfConcurrentTransactions := 20
	errorChannel := make(chan error)
	resultChannel := make(chan TransferTXResult)

	for i := 0; i < numberOfConcurrentTransactions; i++ {
		go func() {
			result, err := testStore.TransferTX(context.Background(), TransferTXParam{
				SenderID:   account1.ID,
				ReceiverID: account2.ID,
				Amount:     utils.RandomNumber(100, 150),
			})

			errorChannel <- err
			resultChannel <- result
		}()
	}

	for i := 0; i < numberOfConcurrentTransactions; i++ {
		err := <-errorChannel
		result := <-resultChannel
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, result.SenderEntry.AccountID, account1.ID)
		require.Equal(t, result.ReceiverEntry.AccountID, account2.ID)
		require.Equal(t, result.Sender.ID, account1.ID)
		require.Equal(t, result.Receiver.ID, account2.ID)
		require.Equal(t, result.SenderEntry.AccountID, result.Sender.ID)
		require.Equal(t, result.ReceiverEntry.AccountID, result.Receiver.ID)
		require.Equal(t, -result.Amount, result.SenderEntry.Amount)
		require.Equal(t, result.Amount, result.ReceiverEntry.Amount)
		require.Equal(t, -result.Transfer.Amount, result.SenderEntry.Amount)
		require.Equal(t, result.Transfer.Amount, result.ReceiverEntry.Amount)
		require.NotZero(t, result.SenderEntry.ID)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.ReceiverEntry.ID)
		require.Equal(t, result.Sender.Balance+result.Receiver.Balance,
			(result.Sender.Balance+result.Amount)+(result.Receiver.Balance-result.Amount))
		_, err = testStore.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)
		_, err = testStore.GetEntry(context.Background(), result.SenderEntry.ID)
		require.NoError(t, err)
		_, err = testStore.GetEntry(context.Background(), result.ReceiverEntry.ID)
		require.NoError(t, err)

		diff1 := account1.Balance - result.Sender.Balance
		diff2 := result.Receiver.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
	}

}

func TestStore_TransferTX_ForDeadlock(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	numberOfConcurrentTransactions := 20
	errorChannel := make(chan error)

	for i := 0; i < numberOfConcurrentTransactions; i++ {
		senderAccount := account1
		receiverAccount := account2

		// Swapping the sender and the receiver for deadlock detection
		if i%2 == 1 {
			senderAccount = account2
			receiverAccount = account1
		}

		go func() {
			_, err := testStore.TransferTX(context.Background(), TransferTXParam{
				SenderID:   senderAccount.ID,
				ReceiverID: receiverAccount.ID,
				Amount:     utils.RandomNumber(100, 150),
			})

			errorChannel <- err
		}()
	}
	// check the final updated balance
	updatedAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
