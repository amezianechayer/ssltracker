package handlers

import (
	"github.com/gofiber/fiber/v2"
)

const localUserKey = "user"

func WithAuthenticatedUser(c *fiber.Ctx) error {
	c.Locals(localUserKey, nil)

	// authentification here..
	// c.Locals(LocalUserKey, )

	//c.Locals(localUserKey, &data.User{ID: 1, Email: "foo@bar.com"})
	return c.Next()
}
