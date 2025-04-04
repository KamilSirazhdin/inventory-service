package handlers

import (
	"github.com/gofiber/fiber/v2"
	"inventory-service/internal/models"
	"inventory-service/internal/repositories"
	"strconv"
)

var repo = repositories.NewItemRepository(db)

func GetItems(c *fiber.Ctx) error {
	items, err := repo.GetAllItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve items",
		})
	}
	return c.JSON(items)
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if err := repo.CreateItem(&item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create item",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

func UpdateItem(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	item.ID = uint(id)
	if err := repo.UpdateItem(&item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update item",
		})
	}
	return c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := repo.DeleteItem(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
