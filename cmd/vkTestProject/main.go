package main

import (
	"context"
	"log"

	"github.com/ReyLegar/vkTestProject/internal/app"
	"github.com/ReyLegar/vkTestProject/internal/config"
	"github.com/ReyLegar/vkTestProject/internal/handler"
	"github.com/ReyLegar/vkTestProject/internal/repository"
	"github.com/ReyLegar/vkTestProject/internal/service"
)

// @title vkTestProject API
// @version 1.0
// @description API для работы с приложением vkTestProject
// @BasePath /
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfig()

	a, err := app.NewApp(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Error!")
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	log.Println("Start app...")

	err = a.Run(handlers)

	if err != nil {
		log.Fatal(err)
	}

}
