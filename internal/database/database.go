package database

import "github.com/go-pg/pg/v10"

func NewDBOptions() *pg.Options {
	return &pg.Options{
		Network:  "tcp",
		Addr:     "localhost:5432",
		Database: "rgb",
		User:     "postgres",
		Password: "1201",
	}
}
