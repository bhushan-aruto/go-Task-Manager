package main

import (
	"github.com/bhushan-aruto/go-task-manager/config"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server"
)

func main() {
	cfg := config.LoadConfig()
	server.Start(cfg)

}
