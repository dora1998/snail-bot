package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE task (
    id varchar(36) NOT NULL PRIMARY KEY,
    body text NOT NULL,
    deadline datetime NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by text
);`

func NewDbInstance(config *DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", config.GetDataSourceName())
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, err
}
