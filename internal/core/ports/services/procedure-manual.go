package ports

import "backend/internal/core/models"

type ProcedureManualService interface {
	CreateProcedureManualService(req models.CreateProcedureManualRequest) error
	GetAllProcedureManualService(limit, offset int) (models.ProcedureManualListResponse, error)
	GetProcedureManualByCategoryService(category string, limit, offset int) (models.ProcedureManualListResponse, error)
	GetProcedureManualByIDService(id int) (*models.ProcedureManualResponse, error)
	UpdateProcedureManualService(id int, req models.UpdateProcedureManualRequest) error
	DeleteProcedureManualService(id int) error
	SearchProcedureManualService(keyword string, limit, offset int) (models.ProcedureManualListResponse, error)
}
