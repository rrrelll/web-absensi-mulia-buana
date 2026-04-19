package middleware

import (
	"os"
	"strings"

	"WebAbsensiMuliaBuana/BackEnd/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON("missing token")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON("invalid token format")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwtlib.ParseWithClaims(tokenString, &jwt.JWTClaims{}, func(t *jwtlib.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON("invalid token")
		}

		claims, ok := token.Claims.(*jwt.JWTClaims)
		if !ok {
			return c.Status(401).JSON("invalid token claims")
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()

	}
}
