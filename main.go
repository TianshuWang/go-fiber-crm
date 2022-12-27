package main

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"go-fiber-crm/repository"
	"go-fiber-crm/service"
)

func main() {
	DBConn := repository.NewDBConn()
	s := service.NewLeadService(repository.NewLeadRepository(DBConn))
	app := fiber.New()
	setupRoutes(app, s)
	err := app.Listen(3000)
	if err != nil {
		return
	}

	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {
			return
		}
	}(DBConn)
}

func setupRoutes(app *fiber.App, s service.LeadService) {
	app.Get("/api/v1/lead/:id", s.GetLead)
	app.Get("/api/v1/leads", s.GetLeads)
	app.Post("/api/v1/lead", s.NewLead)
	app.Delete("/api/v1/lead/:id", s.DeleteLead)
}
