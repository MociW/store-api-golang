package middleware

import (
	"github.com/MociW/store-api-golang/config"
	"github.com/sirupsen/logrus"
)

type MiddlewareConfig struct {
	Logger *logrus.Logger
	Config *config.Config
}

// MiddlewareManager defines methods middleware
type MiddlewareManager struct {
	logger *logrus.Logger
	cfg    *config.Config
}

func NewMiddlewareManager(config *MiddlewareConfig) *MiddlewareManager {
	return &MiddlewareManager{cfg: config.Config, logger: config.Logger}
}
