// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListAccountEntries(ctx context.Context, arg ListAccountEntriesParams) ([]Entry, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListBankEntries(ctx context.Context, arg ListBankEntriesParams) ([]Entry, error)
	ListIncomingTransfers(ctx context.Context, arg ListIncomingTransfersParams) ([]Transfer, error)
	ListOutgoingTransfers(ctx context.Context, arg ListOutgoingTransfersParams) ([]Transfer, error)
	ListTransfersBetween(ctx context.Context, arg ListTransfersBetweenParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)