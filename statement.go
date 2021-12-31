package gopostgres

import (
	"context"
	"database/sql"
	"errors"
)

// Statement type is used to implement statement interface
type Statement struct {
	Stmt *sql.Stmt
}

// ExecContext executes a prepared statement
func (st *Statement) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	data, err := st.Stmt.ExecContext(ctx, args)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return data, nil
}

// QueryContext executes a prepared statement and returns sql.Rows type
func (st *Statement) QueryContext(ctx context.Context, args ...interface{}) (RowsInterface, error) {
	if args == nil {
		return st.Stmt.QueryContext(ctx)
	}

	result, err := st.Stmt.QueryContext(ctx, args)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}

// Close closes the statement
func (st *Statement) Close() error {
	return st.Stmt.Close()
}
