package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type RefreshTokenModel struct {
	RefreshToken string `json:"refresh_token"`
}

func VerifyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte("refresh_token"), nil
}

func (h *Handlers) RefreshToken(c *fiber.Ctx) error {
	var refresh RefreshTokenModel

	if err := c.BodyParser(&refresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := jwt.Parse(refresh.RefreshToken, VerifyToken)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Unauthorized",
			"message": err.Error(),
		})
	}

	if claims_req, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		newToken := jwt.New(jwt.SigningMethodHS256)
		claims := newToken.Claims.(jwt.MapClaims)
		claims["sub"] = claims_req["sub"]
		claims["name"] = "nutthanonn"
		claims["exp"] = time.Now().Add(time.Minute * 3).Unix()
		claims["iat"] = time.Now().Unix()

		t, err := newToken.SignedString([]byte("access_token"))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": t,
		})
	}

	return nil
}
