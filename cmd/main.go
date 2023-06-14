package main

import (
	"context"
	"log"

	"taskman4"
	"taskman4/configs"
	"taskman4/logger"
	"taskman4/pkg/handler"
	"taskman4/pkg/repository"
	"taskman4/pkg/service"

)

func main() {
	configs.PutAdditionalSettings()
	logger.Init()

	logger.Info.Println("taskman started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	repos := repository.NewRepository(cfg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(taskman4.Server)

	if err := srv.Run(ctx, "8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}

	logger.Error.Println("taskman exited")
}
