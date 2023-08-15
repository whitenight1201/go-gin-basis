package server

import (
	"rgb/internal/conf"
	"rgb/internal/database"
	"rgb/internal/store"
)

func Start(cfg conf.Config) {
	jwtSetup(cfg)

	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run("localhost:8080")
}
