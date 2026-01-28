package repositories

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

func (r *UserRepositoryDB) GetEmployeeByFullNameEn(fullNameEn string) (*domains.EmployeeView, error) {
	var emp domains.EmployeeView
	if err := r.db.
		Where("UHR_FullNameEn = ?", fullNameEn).
		First(&emp).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *UserRepositoryDB) GetEmployeeByEmpCode(empCode string) (*domains.EmployeeView, error) {
	var emp domains.EmployeeView
	if err := r.db.
		Where("UHR_EmpCode = ?", empCode).
		First(&emp).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *UserRepositoryDB) FindEmployeeByAccount(account string) (*domains.EmployeeView, error) {
	fmt.Println("account : ", account)
	var user domains.EmployeeView

	if err := r.db.Debug().Where("AD_UserLogon = ?", account).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryDB) GetEmployees() ([]domains.EmployeeView, error) {
	var employees []domains.EmployeeView

	query := `
		SELECT *
		FROM PSTH_HR_SERVICE.dbo.V_HRS_UsersHrService
		WHERE UHR_StatusToUse = @status
	`

	if err := r.db.Raw(
		query,
		sql.Named("status", "ENABLE"),
	).Scan(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}
