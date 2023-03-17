package app

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	errChatNotFound  = errors.New("chat not found")
	errMissingNodeID = errors.New("missing node ID")
)

var FiberErrorHandler = func(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	switch err {
	case errChatNotFound:
		return c.Status(http.StatusNotFound).SendString(err.Error())
	case errMissingNodeID:
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	default:
		return c.Status(http.StatusInternalServerError).SendString("something went wrong")
	}

}
