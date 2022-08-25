package main

import (
	"context"
	"io/ioutil"
	"log"
	"os/signal"
	"syscall"

	"gopkg.in/yaml.v3"

	"github.com/arskydev/weatherman/pkg/db"
	"github.com/arskydev/weatherman/pkg/repository"
	"github.com/arskydev/weatherman/pkg/server"
	"github.com/arskydev/weatherman/pkg/service"
	"github.com/arskydev/weatherman/pkg/web/handlers"
)

type AppConfig struct {
	ConfPath      string `yaml:"CONF_PATH"` // I know it's not an intuitive way to define consts names this way, but it's the Go way.
	PASS_ENV_NAME string `yaml:"PASS_ENV_NAME"`
	APP_PORT      string `yaml:"APP_PORT"`
}

const (
	AppConfPath = "config/app_config.yaml" // move to env variable and set a default value to a static one.
)

func main() {
	appConfig, err := newAppConfig(AppConfPath) // passing a const as param to local func is a bad style
	if err != nil {
		log.Fatal("Error while gathering app config", err)
	}

	pgCfg, err := db.NewPGConfig(appConfig.ConfPath, appConfig.PASS_ENV_NAME)

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

	service := service.NewService(repo)
	handler := handlers.NewHandler(service)
	server := server.NewServer(handler.InitRoutes())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := server.Run(ctx, appConfig.APP_PORT); err != nil {
		log.Printf("Error raised while server run:\n%s", err.Error())
	}

	log.Println("Closing db connect...")
	if err := pgDB.Close(); err != nil {
		log.Fatalf("Error raised while closing db connect:\n%s", err.Error())
	}
	log.Println("Closing db connect... Done")

}

//this func should not be called outside of main func.
// Still it returns an exported object, perhaps should be moved to a separate package, plus using const as param name.
// It makes sense, but this approach is OK too
func newAppConfig(confPath string) (*AppConfig, error) {
	yamlFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return nil, err
	}

	conf := &AppConfig{}
	err = yaml.Unmarshal(yamlFile, conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
