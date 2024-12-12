package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	AWS      AwsConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SSL          bool
	JWTSecretKey string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	NameDB   string
}

type LoggerConfig struct {
	Level       string
	Caller      bool
	Encoding    string
	Development bool
}

type AwsConfig struct {
	Endpoint       string
	MiniEndpoint   string
	MinioAccessKey string
	MinioSecretKey string
	UseSSL         bool
}

func NewAppConfig() (*Config, error) {

	filename := "./config.yaml"

	config := viper.New()
	config.SetConfigFile(filename)
	config.AddConfigPath(".")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file '%s' not found", filename)
		}
		return nil, fmt.Errorf("error reading config file, %v", err)
	}

	cfg := new(Config)
	if err := config.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
