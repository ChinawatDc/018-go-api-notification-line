package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/ChinawatDc/018-go-api-notification-line/internal/config"
	"github.com/ChinawatDc/018-go-api-notification-line/internal/httpserver"
	"github.com/ChinawatDc/018-go-api-notification-line/internal/lineoa"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	lineClient := lineoa.NewClient(lineoa.ClientOptions{
		AccessToken: cfg.LineChannelAccessToken,
	})
	lineSvc := lineoa.NewService(lineoa.ServiceOptions{
		Client:  lineClient,
		AdminTo: cfg.LineAdminTo,
	})

	h := httpserver.NewHandlers(cfg.LineChannelSecret, lineSvc, lineClient, cfg.LineAdminTo)
	r := httpserver.NewRouter(h)

	log.Printf("server running :%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
