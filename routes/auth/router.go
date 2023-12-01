package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app fiber.Router) {
	app.Post("/:cashierId/login", Login)
	app.Get("/:cashierId/logout", Logout)
	app.Post("/cashierId/passcode", Passcode)
}
