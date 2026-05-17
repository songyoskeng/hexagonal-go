package main

import (
	"fmt"
	"hexagonalgo/order/adapters"
	"hexagonalgo/order/core"
	"hexagonalgo/taxcalculator"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("---Tax calculator---")
	taxRepo := taxcalculator.NewTaxRateRepository()
	taxApp := taxcalculator.NewTaxCalculator(taxRepo)

	taxResult := taxApp.CalculateTax(2000)
	fmt.Printf("result: %.2f\n", taxResult)

	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8080")
}
