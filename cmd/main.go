package main

import (
	"log"

	"github.com/MociW/store-api-golang/config"
	"github.com/MociW/store-api-golang/internal/server"
	"github.com/MociW/store-api-golang/pkg/database/aws"
	"github.com/MociW/store-api-golang/pkg/database/postgres"
<<<<<<< HEAD
	"github.com/MociW/store-api-golang/pkg/logger"
=======
	"github.com/MociW/store-api-golang/pkg/database/redis"
	"github.com/MociW/store-api-golang/pkg/email"
	"github.com/gofiber/fiber/v2"
>>>>>>> b5db15a8bb084ecb08d3cfbd59e7d88d79375b51
)

// @version		1.0
// @title			Store API
// @description	API for managing store data
// @BasePath		/api/v1
// @host			localhost:3000
// @schemes		http https
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

<<<<<<< HEAD
	loggerService := logger.NewLogger(cfg)
	if loggerService == nil {
		log.Fatal("Logger is empty")
	}
=======
	redisClient := redis.NewRedis(cfg)

	mailClient := email.NewEmailService(cfg)

	app := fiber.New()
>>>>>>> b5db15a8bb084ecb08d3cfbd59e7d88d79375b51

	s := server.NewServeConfig(&server.ServeConfig{
		Cfg:       cfg,
		Db:        psql,
		AwsClient: awsClient,
<<<<<<< HEAD
		Logger:    loggerService,
=======
		Redis:     redisClient,
		Mail:      mailClient,
>>>>>>> b5db15a8bb084ecb08d3cfbd59e7d88d79375b51
	})

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
