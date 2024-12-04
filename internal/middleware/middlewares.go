package middleware

import "github.com/MociW/store-api-golang/pkg/config"

type MiddlewareConfig struct {
	Config *config.Config
}

type MiddlewareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(config *MiddlewareConfig) *MiddlewareManager {
	return &MiddlewareManager{cfg: config.Config}
}
