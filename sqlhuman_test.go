package main

import (
	"database/sql"
	"net/http"
	"testing"
)

func TestSQLHuman_Add(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlhuman := &SQLHuman{
				db: tt.fields.db,
			}
			sqlhuman.Add(tt.args.w, tt.args.r)
		})
	}
}
