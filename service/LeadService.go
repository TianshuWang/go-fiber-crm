package service

import (
	"github.com/gofiber/fiber"
	"go-fiber-crm/model"
	"go-fiber-crm/repository"
	"net/http"
)

type LeadService interface {
	GetLead(c *fiber.Ctx)
	GetLeads(c *fiber.Ctx)
	NewLead(c *fiber.Ctx)
	DeleteLead(c *fiber.Ctx)
}

type LeadServiceImpl struct {
	repo repository.LeadRepository
}

func NewLeadService() LeadService {
	return &LeadServiceImpl{
		repo: repository.NewLeadRepository(),
	}
}

func (s *LeadServiceImpl) GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	lead := s.repo.FindLead(id)
	if lead.Name == "" {
		c.Status(http.StatusInternalServerError).Send("No lead found with id")
		return
	}
	err := c.JSON(lead)
	if err != nil {
		return
	}
}

func (s *LeadServiceImpl) GetLeads(c *fiber.Ctx) {
	leads := s.repo.FindLeads()
	err := c.JSON(leads)
	if err != nil {
		return
	}
}

func (s *LeadServiceImpl) NewLead(c *fiber.Ctx) {
	lead := new(model.Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(http.StatusServiceUnavailable).Send(err)
		return
	}
	s.repo.CreateLead(lead)
	err := c.JSON(lead)
	if err != nil {
		return
	}
}

func (s *LeadServiceImpl) DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	lead := s.repo.FindLead(id)
	if lead.Name == "" {
		c.Status(http.StatusInternalServerError).Send("No lead found with id")
		return
	}
	s.repo.DeleteLead(&lead)
	c.Send("Lead successfully deleted")
}
