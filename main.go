package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"go-fiber-crm/model"
	"go-fiber-crm/repository"
	"go-fiber-crm/service"
)

var (
	dialect = "sqlite3"
	dbName  = "leads.db"
)

func main() {
	r := repository.NewLeadRepository(initDBConn())
	s := service.NewLeadService(r)
	app := fiber.New()
	setupRoutes(app, s)
	err := app.Listen(3000)
	if err != nil {
		return
	}
}

func initDBConn() *gorm.DB {
	DBConn, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Printf("Connection opened to database")
	DBConn.AutoMigrate(&model.Lead{})
	fmt.Printf("Database migrated")
	return DBConn
}

func setupRoutes(app *fiber.App, s service.LeadService) {
	app.Get("/api/v1/lead/:id", s.GetLead)
	app.Get("/api/v1/leads", s.GetLeads)
	app.Post("/api/v1/lead", s.NewLead)
	app.Delete("/api/v1/lead/:id", s.DeleteLead)
}
