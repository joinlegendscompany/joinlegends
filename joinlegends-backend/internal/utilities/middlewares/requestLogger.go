package middlewares

import (
	"bytes"
	"encoding/json"
	"go-backend-stream/internal/utilities/logger"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(ignorePaths []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		if slices.Contains(ignorePaths, path) {
			return c.Next()
		}

		start := time.Now()

		headers := make(map[string]string)
		c.Request().Header.VisitAll(func(key, value []byte) {
			headers[string(key)] = string(value)
		})

		bodyBytes := c.Body()
		var compactBody bytes.Buffer
		if len(bodyBytes) > 0 {
			if err := json.Compact(&compactBody, bodyBytes); err == nil {
				bodyBytes = compactBody.Bytes()
			}
		}

		headersJSON, _ := json.Marshal(headers)
		logger.Info.Printf("📩 REQUEST_IN: method=%s path=%s headers=%s body=%s",
			c.Method(), c.Path(), string(headersJSON), string(bodyBytes))

		err := c.Next()
		if err != nil {
			return err
		}

		latency := time.Since(start)
		status := c.Response().StatusCode()

		logMsg := map[string]interface{}{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     status,
			"latency_ms": latency.Milliseconds(),
			"client_ip":  c.IP(),
			"user_agent": string(c.Request().Header.UserAgent()),
			"error":      err,
		}

		switch {
		case status >= 500:
			logger.Error.Printf("ROUTE_LOG: %+v", logMsg)
		case status >= 400:
			logger.Warn.Printf("ROUTE_LOG: %+v", logMsg)
		default:
			logger.Info.Printf("ROUTE_LOG: %+v", logMsg)
		}

		return nil
	}
}
