package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MociW/store-api-golang/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	ctxTimeout = 5
	certFile   = "./certs/server.crt"
	keyFile    = "./certs/server.key"
)

type ServeConfig struct {
	Cfg       *config.Config
	Db        *gorm.DB
	AwsClient *minio.Client
}

type Server struct {
	cfg       *config.Config
	db        *gorm.DB
	awsClient *minio.Client
	app       *fiber.App
}

func NewServeConfig(config *ServeConfig) *Server {
	return &Server{
		cfg:       config.Cfg,
		db:        config.Db,
		awsClient: config.AwsClient,
	}
}

func (s *Server) Run() error {
	s.app = fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(s.cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.cfg.Server.WriteTimeout) * time.Second,
	})

	if err := s.Boostrap(); err != nil {
		return errors.Wrap(err, "Server.Run.Bootstrap")
	}

	addr := fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	if s.cfg.Server.SSL {
		serverError := make(chan error)

		go func() {
			serverError <- s.app.ListenTLS(addr, certFile, keyFile)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		select {
		case err := <-serverError:
			log.Fatalf("Failed to start TLS server: %v", err)
		case <-quit:
			if err := s.app.Shutdown(); err != nil {
				log.Fatalf("Error gracefully shutting down server: %v", err)
			}
			log.Println("Server exited properly")
			return nil
		}

	} else {
		serverError := make(chan error)

		go func() {
			serverError <- s.app.Listen(addr)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		select {
		case err := <-serverError:
			log.Fatalf("Failed to start TLS server: %v", err)
		case <-quit:
			if err := s.app.Shutdown(); err != nil {
				log.Fatalf("Error gracefully shutting down server: %v", err)
			}
			log.Println("Server exited properly")
			return nil
		}
	}

	return nil
}
