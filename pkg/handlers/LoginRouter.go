package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nutthanonn/pkg/models"
)

func createToken() (models.Token, error) {
	var msgToken models.Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	user_id := utils.UUID()
	claims["sub"] = user_id
	claims["name"] = "Nutthanonn"
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()
	claims["iat"] = time.Now().Unix()

	t, err := token.SignedString([]byte("access_token"))

	if err != nil {
		return msgToken, err
	}

	msgToken.AccessToken = t

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = user_id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt, err := token.SignedString([]byte("refresh_token"))

	if err != nil {
		return msgToken, err
	}

	msgToken.RefreshToken = rt
	return msgToken, nil
}

func (h *Handlers) LoginRouter(c *fiber.Ctx) error {
	var users models.Users

	if err := c.BodyParser(&users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if users.Username == "nutthanonn" && users.Password == "1234" {
		var token models.Token
		token, err := createToken()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "bad credentials",
	})
}
