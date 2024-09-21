package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/suniln7/go-fiber-crm/database"
	"github.com/suniln7/go-fiber-crm/lead"
)

// function to set routes with app.METHOD("route", function)
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// function to connect to database
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to databse")
	}
	fmt.Println("Connection opened to databse")
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
