package store

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

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

func dbError(_err interface{}) error {
	if _err == nil {
		return nil
	}
	switch _err.(type) {
	case pg.Error:
		err := _err.(pg.Error)
		switch err.Field(82) {
		case "_bt_check_unique":
			return errors.New(extractColumnName(err.Field(110)) + " already exists.")
		}
	case error:
		err := _err.(error)
		switch err.Error() {
		case "pg: no rows in result set":
			return errors.New("not found")
		}
		return err
	}
	return errors.New(fmt.Sprint(_err))
}

func extractColumnName(text string) string {
	reg := regexp.MustCompile(`.+_(.+)_.+`)
	if reg.MatchString(text) {
		return strings.Title(reg.FindStringSubmatch(text)[1])
	}
	return "Unknown"
}
