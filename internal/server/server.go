package server

import (
	"fmt"
	"log"

	"github.com/MociW/store-api-golang/pkg/config"
	"github.com/MociW/store-api-golang/pkg/email"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServeConfig struct {
	App       *fiber.App
	Cfg       *config.Config
	Db        *gorm.DB
	AwsClient *minio.Client
	Redis     *redis.Client
	Mail      email.EmailService
}

type Server struct {
	app       *fiber.App
	cfg       *config.Config
	db        *gorm.DB
	awsClient *minio.Client
	redis     *redis.Client
	mail      email.EmailService
}

func NewServeConfig(config *ServeConfig) *Server {
	return &Server{
		app:       config.App,
		cfg:       config.Cfg,
		db:        config.Db,
		awsClient: config.AwsClient,
		redis:     config.Redis,
		mail:      config.Mail,
	}
}

func (s *Server) Run() error {

	if err := s.Bootstrap(); err != nil {
		return errors.Wrap(err, "Server.Run.Bootstrap")
	}

	err := s.app.ListenTLS(fmt.Sprintf(":%d", s.cfg.Server.Port), "certs/server+3.pem", "certs/server+3-key.pem")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return nil
}
