package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (m *MiddlewareManager) LoggerMidddleware(c *fiber.Ctx) error {
	if m.logger == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Logger is not initialized")
	}

	start := time.Now()

	err := c.Next()

	latency := time.Since(start)
	m.logger.WithFields(logrus.Fields{
		"latency":   latency,
		"method":    c.Method(),
		"uri":       c.Path(),
		"Client-IP": c.IP(),
		"Status":    c.Response().StatusCode(),
	}).Info("HTTP Request completed")

	return err
}
