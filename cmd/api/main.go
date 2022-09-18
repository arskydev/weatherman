package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/arskydev/weatherman/internal/coordinates"
	"github.com/arskydev/weatherman/internal/weather"
	"github.com/arskydev/weatherman/pkg/config"
	"github.com/arskydev/weatherman/pkg/db"
	"github.com/arskydev/weatherman/pkg/repository"
	"github.com/arskydev/weatherman/pkg/server"
	"github.com/arskydev/weatherman/pkg/service"
	"github.com/arskydev/weatherman/pkg/web/handlers"
)

func main() {
	var (
		appConfigPath = "config/app_config.yaml"
		ipGeoKey      = os.Getenv("IPGEO_API_KEY")
		weatherApiKey = os.Getenv("WEATHER_API_KEY")
	)

	appConfig, err := config.NewAppConfig(appConfigPath)
	if err != nil {
		log.Fatal("Error while gathering app config", err)
	}

	pgCfg, err := db.NewPGConfig(appConfig.ConfPath, appConfig.PassEnvName)

	if err != nil {
		log.Fatal("Error while initiating db config:", err)
	}

	pgDB, err := db.NewConnect(pgCfg)

	if err != nil {
		log.Fatal("Error while connecting to DB:", err)
	}

	repo, err := repository.NewRepository(pgDB)
	if err != nil {
		log.Fatal("Error while initiating repository:", err)
	}

	coordinator := coordinates.New(ipGeoKey)
	weatherer := weather.New(coordinator, weatherApiKey)

	service := service.NewService(repo)
	handler := handlers.NewHandler(service, weatherer)
	server := server.NewServer(handler.InitRoutes())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := server.Run(ctx, appConfig.AppPort); err != nil {
		log.Printf("Error raised while server run:\n%s", err.Error())
	}

	log.Println("Closing db connect...")
	if err := pgDB.Close(); err != nil {
		log.Fatalf("Error raised while closing db connect:\n%s", err.Error())
	}
	log.Println("Closing db connect... Done")

}
