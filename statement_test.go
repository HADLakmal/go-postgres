package gopostgres_test

import (
	"bitbucket.org/mybudget-dev/go-postgres"
	"context"
	"testing"
)

func TestStatement_ExecContext(t *testing.T) {
	newDB := gopostgres.DatabaseReporter(&gopostgres.DBPool{DB: newTestDB(t, "people")})
	defer newDB.Close()

	tests := map[string]struct {
		query   string
		argName string
		wantErr bool
	}{
		"norma-query": {
			query:   "INSERT|people|name=?",
			argName: "posgres",
		},
		"failed-query": {
			query:   "INSERT|people|name=?,age=?",
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			tx, err := newDB.BeginTransaction(ctx, nil)
			if err != nil {
				t.Fatal(err)
			}

			p, err := tx.PrepareContext(ctx, tt.query)
			if err != nil {
				t.Fatal(err)
			}

			_, err = p.ExecContext(ctx, tt.argName)
			if err != nil && !tt.wantErr {
				t.Error(err)
			}

			err = tx.Commit()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestStatement_QueryContext(t *testing.T) {
	newDB := gopostgres.DatabaseReporter(&gopostgres.DBPool{DB: newTestDB(t, "people")})
	defer newDB.Close()

	tests := map[string]struct {
		query   string
		wantErr bool
	}{
		"norma-query": {
			query:   "SELECT|people|name|",
		},
		"failed-query": {
			query:   "SELECT|people|name=?|",
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			p, err := newDB.Prepare(ctx, tt.query)
			if err != nil {
				t.Fatal(err)
			}

			_, err = p.QueryContext(ctx)
			if err != nil && !tt.wantErr {
				t.Error(err)
			}

			err = p.Close()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
