package domain_upload

import (
	"context"
	"go-backend-stream/internal/infrastructure/redisclient"
	"go-backend-stream/internal/utilities/logger"

	"github.com/gofiber/fiber/v2"
)

func VideoUploadHandler(c *fiber.Ctx) error {
	logger.Info.Printf("video called successful")
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// caminho para a pasta compartilhada
	sharedDir := "/shared/videos"

	// salva o arquivo lá
	if err := c.SaveFile(file, sharedDir+"/"+file.Filename); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	redisCtx := context.Background()
	rdb := redisclient.New()

	msg := redisclient.QueueMessage{
		ID:   "uuid-123",
		Type: "view.progress",
		Payload: map[string]string{
			"file": file.Filename,
		},
	}

	if err := redisclient.Send(redisCtx, rdb, msg); err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"filename": file.Filename,
		"size":     file.Size,
	})
}
