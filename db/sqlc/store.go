package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(queries *Queries, db *sql.DB) (*Store, error) {
	return &Store{queries, db}, nil
}

type TransferTXParam struct {
	SenderID   int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
	Amount     int64 `json:"amount"`
}

type TransferTXResult struct {
	Sender        Account  `json:"sender"`
	Receiver      Account  `json:"receiver"`
	Amount        int64    `json:"amount"`
	SenderEntry   Entry    `json:"senderEntry"`
	ReceiverEntry Entry    `json:"receiverEntry"`
	Transfer      Transfer `json:"transfer"`
}

func (store *Store) executeTransaction(ctx context.Context, fn func(q *Queries) error) error {
	transaction, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(transaction)
	err = fn(q)
	if err != nil {
		if rollBackErr := transaction.Rollback(); rollBackErr != nil {
			finalError := fmt.Errorf("error:%s, rollback error: %s", err, rollBackErr)
			return finalError
		}
		return err
	}
	return transaction.Commit()
}

func (store *Store) TransferTX(ctx context.Context, params TransferTXParam) (TransferTXResult, error) {
	var result TransferTXResult
	var err error
	err = store.executeTransaction(ctx, func(q *Queries) error {

		// Create the transfer record
		transfer, err := q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: params.SenderID,
			ToAccountID:   params.ReceiverID,
			Amount:        params.Amount,
		})
		if err != nil {
			return err
		}
		result.Transfer = transfer

		// Create the sender entry
		entry, err := q.CreateEntry(ctx, CreateEntryParams{
			AccountID: params.SenderID,
			Amount:    -1 * params.Amount,
		})

		if err != nil {
			return err
		}
		result.SenderEntry = entry

		// Create the receiver entry
		entry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: params.ReceiverID,
			Amount:    params.Amount,
		})

		if err != nil {
			return err
		}
		result.ReceiverEntry = entry
		result.Amount = transfer.Amount

		// This is to make sure the lock is obtained in same order always to avoid deadlock
		if params.SenderID < params.ReceiverID {
			// Edit the sender and the receiver balance and return their accounts
			sender, err := q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      params.SenderID,
				Balance: -params.Amount,
			})
			if err != nil {
				return err
			}

			receiver, err := q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      params.ReceiverID,
				Balance: params.Amount,
			})
			if err != nil {
				return err
			}
			result.Sender = sender
			result.Receiver = receiver
		} else {
			receiver, err := q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      params.ReceiverID,
				Balance: params.Amount,
			})
			if err != nil {
				return err
			}

			// Edit the sender and the receiver balance and return their accounts
			sender, err := q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      params.SenderID,
				Balance: -params.Amount,
			})
			if err != nil {
				return err
			}
			result.Sender = sender
			result.Receiver = receiver
		}

		return nil
	})
	if err != nil {
		return TransferTXResult{}, err
	}

	return result, nil
}
