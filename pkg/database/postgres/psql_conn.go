package postgres

import (
	"fmt"
	"time"

	"github.com/MociW/store-api-golang/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	username := config.Postgres.User
	password := config.Postgres.Password
	host := config.Postgres.Host
	port := config.Postgres.Port
	database := config.Postgres.NameDB

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, port, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}

// migrate -database "postgres://postgres:postgres@localhost:5432/negodb?sslmode=disable" -path db/migrations up
// migrate -database "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable" -path db/migrations down
