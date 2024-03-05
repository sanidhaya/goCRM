package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/sanidhaya/goCRM/database"
	"github.com/sanidhaya/goCRM/lead"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLeads)
	app.Delete("api/v1/lead:id", lead.DeleteLeads)
}

func initDatabase() {
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to databse")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
