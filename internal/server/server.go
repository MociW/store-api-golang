package server

import (
	"fmt"
	"log"

	"github.com/MociW/store-api-golang/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ServeConfig struct {
	App       *fiber.App
	Cfg       *config.Config
	Db        *gorm.DB
	AwsClient *minio.Client
}

type Server struct {
	app       *fiber.App
	cfg       *config.Config
	db        *gorm.DB
	awsClient *minio.Client
}

func NewServeConfig(config *ServeConfig) *Server {
	return &Server{
		app:       config.App,
		cfg:       config.Cfg,
		db:        config.Db,
		awsClient: config.AwsClient,
	}
}

func (s *Server) Run() error {

	if err := s.Boostrap(); err != nil {
		return errors.Wrap(err, "Server.Run.Bootstrap")
	}

	// cert := &tls.Config{
	// 	MinVersion:   tls.VersionTLS12,
	// 	Certificates: []tls.Certificate{},
	// }

	// err := s.app.Listen(fmt.Sprintf(":%d", s.cfg.Server.Port))
	// if err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }

	err := s.app.ListenTLS(fmt.Sprintf(":%d", s.cfg.Server.Port), "certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Server.Port))

	// cer, _ := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")

	// ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})

	// s.app.Listener(ln)

	return nil
}
