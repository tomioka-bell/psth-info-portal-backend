package ports

import (
	"backend/internal/core/domains"
	"backend/internal/core/models"
)

type UserRepository interface {
	CreateUserRepository(User *domains.User) error
	FindByUsername(username string) (*domains.User, error)
	GetUserByID(userID string) (domains.User, error)
	GetAllUser() ([]domains.User, error)
	UpdateUserWithMap(userID string, updates map[string]interface{}) error
	GetUserCount() (int64, error)
	GetEmployeeByFullNameEn(fullNameEn string) (*domains.EmployeeView, error)

	// ====================== Employee View ===================================
	FindEmployeeByAccount(account string) (*domains.EmployeeView, error)
	GetEmployeeByEmpCode(empCode string) (*domains.EmployeeView, error)
	GetEmployees() ([]domains.EmployeeView, error)

	// ====================== Employee ===================================
	SaveOrUpdateEmployee(emp *models.PSEmployee) error
}
