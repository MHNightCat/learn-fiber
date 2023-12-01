package cashier

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app fiber.Router) {
	app.Post("/", createCashier)
	app.Get("/", cashierList)
	app.Get("/:cashierId", getCashierDetails)
	app.Delete("/:cashierId", deleteCashier)
	app.Put("/:cashierId", updateCashier)
}
