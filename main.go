package main

import (
	"chatnodes/app"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setup repository
	repo := app.NewMemChatRepository()

	// Setup service
	svc := app.NewChatServiceImpl(repo)

	// Setup controller
	cont := app.NewChatController(svc)

	// Setup fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: app.FiberErrorHandler,
	})

	// Setup routing
	cont.Route(app)

	// Start app
	app.Listen(":3000")
}
