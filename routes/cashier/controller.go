package cashier

import (
	db "learn-fiber/config"
	models "learn-fiber/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func createCashier(c *fiber.Ctx) error {
	var data map[string]string
	// get body data and store it in data
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier Name is required",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier passcode is required",
			})
	}

	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	db.DB.Create(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier added successfully",
		"data":    cashier,
	})
}

func cashierList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier list api",
			"data":    cashier,
		})
}

func updateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id=?", cashierId)
	
	if cashier.Name == ""{
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			})
	}

	var updateCashier models.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			})
	}

	cashier.Name = updateCashier.Name
	db.DB.Save(&cashier)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "success",
			"data": cashier,
		})
}

func getCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Select("id,name,created_at,update_at").Where("id=?", cashierId).First(&cashier)
	
	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createAt"] = cashier.CreatedAt
	cashierData["updateAt"] = cashier.UpdateAt

	if cashier.Id == 0{
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
				"error": map[string]interface{}{},
			})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashierData,
	})
}

func deleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Where("id = ?",cashierId).First(&cashier)
	if cashier.Id == 0{
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			})
	}

	db.DB.Where("id = ?", cashierId).Delete(&cashier)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier deleted successfully",
		})
}
