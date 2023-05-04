package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"main/clock"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Preparer interface {
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Queryer interface {
	Preparer
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Queryer = (*sqlx.DB)(nil)
	_ Execer = (*sqlx.DB)(nil)
	_ Execer = (*sqlx.Tx)(nil)
)

type Repository struct {
	Clocker clock.Clocker
}

const (
	ErrCodeMYSQLDuplicateEntry = 1062
)

var (
	ErrAlreadyEntry = errors.New("duplicate entry")
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	db, err := sql.Open("mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?parseTime=true",
				cfg.DBUser, cfg.DBPassword,
				cfg.DBHost, cfg.DBPort,
				cfg.DBName,
			),
		)

	if err != nil {
		return nil, nil, err
	}

	ctx, cansel := context.WithTimeout(ctx,2*time.Second)
	defer cansel()

	if err := db.PingContext(ctx); err != nil {
		return nil, func() {_ = db.Close()}, err
	}

	xdb := sqlx.NewDb(db,"mysql")
	return xdb, func() {_ = db.Close()}, nil

}