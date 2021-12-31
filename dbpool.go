package gopostgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

// SSLMode ssl mode
type SSLMode string

// SSL mode types
const (
	Disable    SSLMode = `disable`     // No SSL
	Require    SSLMode = `require`     // Always SSL (skip verification)
	VerifyCA   SSLMode = `verify-ca`   // signed by a trusted CA
	VerifyFull SSLMode = `verify-full` // Always SSL

)

// DBConfig configuration of posgres
type DBConfig struct {
	User              string
	Password          string
	Host              string
	Port              string
	Database          string
	ConnectionTimeout time.Duration
	SSLMode
}

// DBPool inherit behavior of sql
type DBPool struct {
	*sql.DB
}

// NewDBPool posgreSQL database connection
func NewDBPool(conf DBConfig) (DatabaseReporter, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	// SSLMode configure
	if conf.SSLMode == `` {
		connString += fmt.Sprintf(`?sslmode=%s`, Disable)
	} else {
		connString += fmt.Sprintf(`?sslmode=%s`, conf.SSLMode)
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(conf.ConnectionTimeout)

	return &DBPool{db}, nil
}

// BeginTransaction starts a transaction in postgres
func (a *DBPool) BeginTransaction(ctx context.Context, opts *sql.TxOptions) (TransactionInterface, error) {
	tx, err := a.BeginTx(ctx, opts)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &Transaction{
		tx,
	}, nil
}

// Prepare prepares a row statement
func (a *DBPool) Prepare(ctx context.Context, sql string) (StatementInterface, error) {
	statement, err := a.PrepareContext(ctx, sql)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &Statement{
		Stmt: statement,
	}, nil
}

// Stop will close the Postgres adapter releasing connections
func (a *DBPool) Stop() error {
	return a.Close()
}
