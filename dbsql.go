package gopostgres

import (
	"context"
	"database/sql"
)

//go:generate mockgen -destination=./mock_sql.go -package=gopostgres -source=./dbsql.go

type DatabaseReporter interface {
	// Close will close the db reporter
	Close() error
	// Prepare prepares a SQL statement
	Prepare(ctx context.Context, sql string) (StatementInterface, error)

	// BeginTransaction returns a transaction interface that can be used to execute multiple statements
	BeginTransaction(context.Context, *sql.TxOptions) (TransactionInterface, error)
}

// TransactionInterface will be used to execute multiple SQL statements in a reversible way
type TransactionInterface interface {

	// PrepareContext prepares a single SQL statement
	PrepareContext(ctx context.Context, query string) (StatementInterface, error)

	// Commit the transaction
	Commit() error

	// Rollback reverts the changes
	Rollback() error
}

// StatementInterface will be used to generalize SQL statements throughout the application
type StatementInterface interface {
	// ExecContext and SQL statement
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)

	// QueryContext the sql data
	QueryContext(ctx context.Context, args ...interface{}) (RowsInterface, error)

	// Close the statement
	Close() error
}

// RowsInterface will be used to generalize SQL scans throughout the rows
type RowsInterface interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
}
