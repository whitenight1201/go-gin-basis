package main

import (
	"rgb/internal/conf"
	"rgb/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}
