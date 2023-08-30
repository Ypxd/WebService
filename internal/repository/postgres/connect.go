package postgres

import (
	"github.com/Ypxd/WebService/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"time"
)

func Connect(cfg models.DB) (*sqlx.DB, error) {
	var err error

	db, err := sqlx.Open(cfg.DriverName, cfg.ConnString)
	if err != nil {
		log.Println(errors.WithMessage(err, "connect to DB"))
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(cfg.ConnTimeout) * time.Second)
	db.SetMaxIdleConns(cfg.MaxConns)
	db.SetMaxOpenConns(cfg.MaxConns)

	err = db.Ping()
	return db, errors.WithMessage(err, "ping db")
}
