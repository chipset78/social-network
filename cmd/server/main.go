package main

import (
	"log"
	"social-network/internal/app"
	"social-network/internal/config"
	"social-network/internal/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewDatabase(&repository.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	server := app.NewServer(db, cfg)
	log.Printf("Server starting on :%s", cfg.AppPort)
	if err := server.Start(":" + cfg.AppPort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
