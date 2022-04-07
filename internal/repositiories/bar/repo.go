package bar

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Repository interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type repo struct {
	db          Repository
	newRelicApp *newrelic.Application
}

func New(newRelicApp *newrelic.Application, db Repository) *repo {
	return &repo{
		db:          db,
		newRelicApp: newRelicApp,
	}
}

func (r *repo) FindUserByBars(ctx context.Context, bar string) ([]*User, error) {
	txn := r.newRelicApp.StartTransaction("FindUserByBars")
	defer txn.End()

	ctx = newrelic.NewContext(ctx, txn)
	rows, err := r.db.QueryContext(ctx, "select name from users where bar = $1", bar)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	bars := make([]*User, 0)
	for rows.Next() {
		bar := User{}
		err := rows.Scan(&bar.Name)
		if err != nil {
			return nil, err
		}
		bars = append(bars, &bar)
	}
	return bars, nil
}
