package handlers

import (
	"github.com/amezianechayer/ssltracker/data"
	"github.com/gofiber/fiber/v2"
)

func getAuthenticatedUser(c *fiber.Ctx) *data.User {
	value := c.Locals(localUserKey)
	if user, ok := value.(*data.User); ok {
		return user
	}
	return nil
}
