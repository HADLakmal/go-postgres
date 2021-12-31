package gopostgres

import (
	"context"
	"database/sql"
	"errors"
)

// Transaction implements transaction interface
type Transaction struct {
	txn *sql.Tx
}

// Commit commits SQL transaction
func (t *Transaction) Commit() error {
	err := t.txn.Commit()
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// PrepareContext prepares SQL statement in a transaction
func (t *Transaction) PrepareContext(ctx context.Context, query string) (StatementInterface, error) {
	stmt, err := t.txn.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &Statement{
		Stmt: stmt,
	}, nil
}

// Rollback reverts the transaction
func (t *Transaction) Rollback() error {
	return t.txn.Rollback()
}
