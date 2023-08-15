package server

import (
	"rgb/internal/database"
	"rgb/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()

	// Start listening and serving requests
	router.Run("localhost:8080")
}
