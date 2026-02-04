package ports

import (
	"backend/internal/core/domains"
	"backend/internal/core/models"
)

type CustomerManualService interface {
	CreateCustomerManualService(req models.CreateCustomerManualRequest) error
	CreateCustomerManualWithResponse(req models.CreateCustomerManualRequest) (*models.CustomerManualResponse, error)
	GetAllCustomerManualService(limit, offset int) (models.CustomerManualListResponse, error)
	GetAllCustomerManualsTree() ([]domains.CustomerManualMenu, error)
	GetCustomerManualByCategoryService(category string, limit, offset int) (models.CustomerManualListResponse, error)
	GetCustomerManualByIDService(id int) (*models.CustomerManualResponse, error)
	UpdateCustomerManualService(id int, req models.UpdateCustomerManualRequest) error
	DeleteCustomerManualService(id int) error
	SearchCustomerManualService(keyword string, limit, offset int) (models.CustomerManualListResponse, error)
}
