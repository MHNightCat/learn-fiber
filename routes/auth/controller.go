package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	db "learn-fiber/config"
	models "learn-fiber/model"
	"os"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid post request",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode is required",
				"error":   map[string]interface{}{},
			})
	}

	var cashier models.Cashier

	db.DB.Where("id = ?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "cashier not found",
				"error":   map[string]interface{}{},
			})
	}

	if cashier.Passcode != data["passcode"] {
		return c.Status(401).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode not match",
				"error":   map[string]interface{}{},
			})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":  strconv.Itoa(int(cashier.Id)),
		"Expires": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(401).JSON(
			fiber.Map{
				"success": false,
				"message": "Token Expired or Invalied",
			})
	}

	cashierData := make(map[string]interface{})
	cashierData["token"] = tokenString

	cookie := new(fiber.Cookie)
	cookie.Name = "accessToken"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)
	return c.Status((200)).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashierData,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

func Passcode(c *fiber.Ctx) error {
	return c.SendString("Passcode")
}
