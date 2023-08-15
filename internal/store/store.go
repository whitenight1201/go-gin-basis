package store

import (
	"log"

	"github.com/go-pg/pg/v10"
)

// Database connector
var db *pg.DB

func SetDBConnection(dbOpts *pg.Options) {
	if dbOpts == nil {
		log.Panicln("DB options can't be nil")
	} else {
		db = pg.Connect(dbOpts)
	}
}

func GetDBConnection() *pg.DB { return db }
