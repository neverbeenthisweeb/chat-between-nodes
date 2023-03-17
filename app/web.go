package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type chatController struct {
	svc ChatService
}

func NewChatController(svc ChatService) *chatController {
	return &chatController{svc}
}

func (ct *chatController) Route(app *fiber.App) {
	app.Use(logger.New())
	app.Get("/node/:id/chats", ct.GetByNodeID)
}

func (ct *chatController) GetByNodeID(c *fiber.Ctx) error {
	nodeID := c.Params("id")

	chats, err := ct.svc.GetByNodeID(c.Context(), nodeID)
	if err != nil {
		return err
	}

	return c.JSON(chats)
}
