package dao

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SQLHelper struct {
	db *sql.DB
}

func NewSQLHelper(driver, connstr string) (*SQLHelper, error) {
	db, err := sql.Open(driver, connstr)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS article (
			id varchar(128),
			name varchar(128) NOT NULL,
			author varchar(128) NOT NULL,
			brief TEXT NOT NULL,
			content TEXT NOT NULL,
			PRIMARY KEY (id)
			)`)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	logrus.Debug("database initialized")
	return &SQLHelper{db:db}, nil
}

func DBHandler(driver, connstr string) func(handler http.Handler) http.Handler {
	helper, err := NewSQLHelper(driver, connstr)
	if err != nil {
		panic(err)
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "db", helper)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}