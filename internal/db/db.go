package db

import (
	"time"

	"github.com/2026-snapy/recommend-server/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	client *sqlx.DB
}

func NewDB(cfg *config.Config) (*DB, error) {
	db, err := sqlx.Connect("mysql", cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{db}, nil
}