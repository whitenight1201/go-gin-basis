package database

import (
	"rgb/internal/conf"

	"github.com/go-pg/pg/v10"
)

func NewDBOptions(cfg conf.Config) *pg.Options {
	return &pg.Options{
		Network:  "tcp",
		Addr:     cfg.DbHost + ":" + cfg.DbPort,
		Database: cfg.DbName,
		User:     cfg.DbUser,
		Password: cfg.DbPassword,
	}
}
