package database

import (
	"database/sql"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
)

func NewDBPool() *sql.DB {
	db, err := sql.Open("nrpgx", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		panic("failed to connect with configs " + err.Error())
	}

	return db
}
