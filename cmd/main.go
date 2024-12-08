package main

import (
	"log"

	"github.com/MociW/store-api-golang/internal/server"
	"github.com/MociW/store-api-golang/pkg/config"
	"github.com/MociW/store-api-golang/pkg/database/aws"
	"github.com/MociW/store-api-golang/pkg/database/postgres"
	"github.com/gofiber/fiber/v2"
)

//	@version		1.0
//	@title			Store API
//	@description	API for managing store data
//	@BasePath		/api/v1
//	@host			localhost:3000
//	@schemes		http https
func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	psql, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	awsClient, err := aws.NewAWSClient(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	app := fiber.New()

	s := server.NewServeConfig(&server.ServeConfig{
		App:       app,
		Cfg:       cfg,
		Db:        psql,
		AwsClient: awsClient,
	})

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
