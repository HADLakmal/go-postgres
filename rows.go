package gopostgres

import "database/sql"

// Rows type is used to implement scan for rows
type Rows struct {
	*sql.Rows
}

func (r *Rows) Next() bool {
	return r.Next()
}

func (r *Rows) Scan(dest ...interface{}) error {
	return r.Scan(dest)
}

func (r *Rows) Close() error {
	return r.Close()
}
