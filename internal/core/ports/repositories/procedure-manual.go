package ports

import "backend/internal/core/domains"

type ProcedureManualRepository interface {
	CreateProcedureManual(manual *domains.ProcedureManual) error
	GetProcedureManualByID(id int) (*domains.ProcedureManual, error)
	GetAllProcedureManuals(page, pageSize int) ([]domains.ProcedureManual, int64, error)
	GetProcedureManualsByCategory(category string, page, pageSize int) ([]domains.ProcedureManual, int64, error)
	UpdateProcedureManual(id int, updates map[string]interface{}) error
	DeleteProcedureManual(id int) error
	SearchProcedureManuals(keyword string, page, pageSize int) ([]domains.ProcedureManual, int64, error)
}
