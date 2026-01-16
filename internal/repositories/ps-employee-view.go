package repositories

import (
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
	const limit int = 5000

	if err := r.db.
		Where("UHR_StatusToUse <> ?", "DISABLE").
		Limit(limit).
		Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}
