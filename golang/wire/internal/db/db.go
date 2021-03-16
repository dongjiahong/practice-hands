package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"

	"learnwire/internal/config"
)

var Provider = wire.NewSet(New) // 同理

func New(cfg *config.Config) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, nil
}
