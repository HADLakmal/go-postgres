package gopostgres_test

import (
	"bitbucket.org/mybudget-dev/go-postgres"
	"context"
	"testing"
)

func TestDBPool_BeginTransaction(t *testing.T) {
	newDB := gopostgres.DatabaseReporter(&gopostgres.DBPool{DB: newTestDB(t, "begin")})
	defer newDB.Close()

	_, err := newDB.BeginTransaction(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDBPool_Prepare(t *testing.T) {
	newDB := gopostgres.DatabaseReporter(&gopostgres.DBPool{DB: newTestDB(t, "prepare")})
	defer newDB.Close()

	tests := map[string]struct {
		query   string
		wantErr bool
	}{
		"norma-query":  {query: "WAIT|1s|SELECT|people|age,name|"},
		"failed-query": {query: "select error", wantErr: true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := newDB.Prepare(context.Background(), tt.query)
			if err != nil && !tt.wantErr {
				t.Error(err)
			}
		})
	}
}
