package main

import (
	"fmt"
	db "learn-fiber/config"
	auth "learn-fiber/routes/auth"
	"learn-fiber/routes/cashier"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	fmt.Println("Fiber Is Starting... 🚀")
	//connect to database(from /config/dbConnection.go)
	db.Connect()

	app := fiber.New()
		
	// Set logger
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} - | ${status} |${latency}     |   ${method}   | ${path} \n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "local",
	}))

	api_v1 := app.Group("/v1")

	//set auth controller
	auth_group := api_v1.Group("/auth")
	auth.Setup(auth_group)

	//set cashier controller
	cashier_group := api_v1.Group("/cashier")
	cashier.Setup(cashier_group)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, exitone!")
	})

	// app Listen
	app.Listen(":8080")

}
 