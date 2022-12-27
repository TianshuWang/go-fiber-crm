package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"go-fiber-crm/model"
)

var (
	dialect = "sqlite3"
	dbName  = "leads.db"
)

type LeadRepository interface {
	FindLead(id string) model.Lead
	FindLeads() []model.Lead
	CreateLead(lead *model.Lead)
	DeleteLead(lead *model.Lead)
}

type LeadRepositoryImpl struct {
	DBConn *gorm.DB
}

func NewLeadRepository() LeadRepository {
	return &LeadRepositoryImpl{
		DBConn: NewDBConn(),
	}
}

func NewDBConn() *gorm.DB {
	DBConn, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Printf("Connection opened to database")
	DBConn.AutoMigrate(&model.Lead{})
	fmt.Printf("Database migrated")
	return DBConn
}

func (r *LeadRepositoryImpl) FindLead(id string) model.Lead {
	var lead model.Lead
	r.DBConn.Find(&lead, id)
	return lead
}

func (r *LeadRepositoryImpl) FindLeads() []model.Lead {
	var leads []model.Lead
	r.DBConn.Find(&leads)
	return leads
}

func (r *LeadRepositoryImpl) CreateLead(lead *model.Lead) {
	r.DBConn.Create(&lead)
}

func (r *LeadRepositoryImpl) DeleteLead(lead *model.Lead) {
	r.DBConn.Delete(&lead)
}
