package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vTCP-Foundation/observerd/common/settings"
	"github.com/vTCP-Foundation/observerd/core/ec"
)

var (
	db *pgxpool.Pool
)

func DB() *pgxpool.Pool {
	if db == nil {
		ec.InterruptOnError(initDBPool())
		ec.InterruptOnError(EnsureSchema())
	}

	return db
}

func initDBPool() (err error) {
	db, err = pgxpool.Connect(context.Background(), settings.Conf.Database.ConnectionCredentials())
	return
}

func NewTransaction() (tx pgx.Tx, err error) {
	tx, err = DB().Begin(context.Background())
	return
}

func Commit(tx pgx.Tx) (err error) {
	err = tx.Commit(context.Background())
	return
}

func RollbackSafely(tx pgx.Tx) {
	_ = tx.Rollback(context.Background())
}

func Query(tx pgx.Tx, sql string, args ...interface{}) (rows pgx.Rows, err error) {
	rows, err = tx.Query(context.Background(), sql, args...)
	return
}

func QueryRow(tx pgx.Tx, sql string, args ...interface{}) (row pgx.Row) {
	row = tx.QueryRow(context.Background(), sql, args...)
	return
}

func Exec(tx pgx.Tx, sql string) (err error) {
	_, err = tx.Exec(context.Background(), sql)
	return
}
