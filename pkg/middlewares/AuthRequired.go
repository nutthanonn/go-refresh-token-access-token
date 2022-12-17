package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte("access_token"),
		SigningMethod:  "HS256",
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
	})
}
