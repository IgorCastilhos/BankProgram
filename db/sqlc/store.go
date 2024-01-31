package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store define todas as funções para executas queries e transações no banco de dados
type Store interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// SQLStore fornece todas as funções para executas queries SQL e transações
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore cria uma nova store (armazenamento)
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
