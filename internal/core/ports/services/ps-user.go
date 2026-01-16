package ports

import (
	"backend/internal/core/models"
)

type UserService interface {
	CreateUserService(req models.UserResp) error
	SignIn(dto models.LoginEmpResp) (string, error)
	GetProfileByCookieId(userID string) (models.UserReq, error)
	GetAllUserSevice() ([]models.UserReqAll, error)
	UpdateUserWithMapService(userID string, updates map[string]interface{}) error
	GetUserByID(userID string) (models.UserAdminReq, error)
	GetUserCountService() (int64, error)

	// ====================== Employee View ===================================
	SignInEmployee(dto models.LoginEmpResp) (string, error)
	GetEmployeeByEmpCodeService(userID string) (models.EmployeeViewByEmpCodeResp, error)
	GetAllEmployeesService() ([]models.EmployeeViewResp, error)
}
