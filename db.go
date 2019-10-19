package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
)

func NewDbInstance(config *DatabaseConfig) (*sqlx.DB, error) {
	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	db, err := sqlx.Connect("mysql", config.GetDataSourceName())
	if err != nil {
		return nil, err
	}

	_, err = migrate.Exec(db.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}

	return db, err
}
