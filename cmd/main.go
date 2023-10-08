package main

import (
	"log"

	"github.com/anton-okolelov/json-app/internal/config"
	"github.com/anton-okolelov/json-app/internal/database"
	"github.com/anton-okolelov/json-app/internal/httpserver"
	"github.com/anton-okolelov/json-app/internal/service"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := database.InitDB(cfg.DBConf)
	if err != nil {
		log.Fatal("cannot init DB:", err)
	}

	userService := service.NewService(db)
	server := httpserver.New(userService, cfg.Port)

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
