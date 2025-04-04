package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"inventory-service/internal/handlers"
	"inventory-service/internal/repositories"
	"log"
)

var db *gorm.DB

func main() {
	// Подключение к базе данных
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=user dbname=inventory_db password=password port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	// Создание таблиц, если они не существуют
	db.AutoMigrate(&repositories.Item{})

	app := fiber.New()

	// Роуты
	app.Get("/items", handlers.GetItems)
	app.Post("/items", handlers.CreateItem)
	app.Put("/items/:id", handlers.UpdateItem)
	app.Delete("/items/:id", handlers.DeleteItem)

	docker --version
	// Запуск сервера
	app.Listen(":3000")
}
