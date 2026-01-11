package main

import (
	"fmt"
	"log"

	"go-backend-stream/internal/infrastructure/api/routes"
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/infrastructure/redisclient"
	"go-backend-stream/internal/utilities/config"
	"go-backend-stream/internal/utilities/logger"
	"go-backend-stream/internal/utilities/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	logger.Init()
	config.LoadEnv()

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024 * 1024,
	})

	rdb := redisclient.New()
	ctx := redisclient.Ctx()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		defer conn.Close()

		pubsub := rdb.Subscribe(ctx, "videos_updates")
		defer pubsub.Close()

		ch := pubsub.Channel()

		for msg := range ch {
			if err := conn.WriteMessage(
				websocket.TextMessage,
				[]byte(msg.Payload),
			); err != nil {
				logger.Error.Println("WS error:", err)
				break
			}
		}
	}))

	app.Use(middlewares.RequestLogger([]string{"/v1/upload/video"}))

	app.Static("/", "./api-doc")

	logger.Info.Println("connecting in the database")
	db, err := database.Connect()
	if err != nil {
		logger.Error.Printf("error when connect in the database: %s\n", err.Error())
		log.Fatal(err)
	}
	defer db.Close()

	logger.Info.Println("running auto migrations")
	if err := db.AutoMigrate(); err != nil {
		logger.Error.Printf("error when running auto migrations: %s\n", err.Error())
		log.Fatal(err)
	}
	routes.Setup(app, db)

	port := ":8080"
	fmt.Printf("Start the server in http://localhost%s\n", port)
	log.Fatal(app.Listen(port))
}
