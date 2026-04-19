package middleware

import "github.com/gofiber/fiber/v2"

func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role").(string)

		for _, r := range roles {
			if userRole == r {
				return c.Next()
			}
		}

		return c.Status(403).JSON("forbidden")
	}
}
