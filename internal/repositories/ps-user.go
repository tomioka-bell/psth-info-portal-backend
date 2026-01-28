package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	ports "backend/internal/core/ports/repositories"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {
	if err := db.AutoMigrate(&domains.User{}, &domains.PSEmployee{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) CreateUserRepository(User *domains.User) error {
	if err := r.db.Create(User).Error; err != nil {
		fmt.Printf("CreateUserRepository error: %v\n", err)
		return err
	}
	return nil
}

func (r *UserRepositoryDB) FindByUsername(username string) (*domains.User, error) {
	var user domains.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryDB) GetUserByID(userID string) (domains.User, error) {
	var user domains.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (r UserRepositoryDB) GetAllUser() ([]domains.User, error) {
	var reviews []domains.User
	return reviews, r.db.Find(&reviews).Error
}

func (r UserRepositoryDB) UpdateUserWithMap(userID string, updates map[string]interface{}) error {
	return r.db.Model(&domains.User{}).
		Where("user_id = ?", userID).
		Updates(updates).
		Error
}

func (r UserRepositoryDB) GetUserCount() (int64, error) {
	var count int64
	return count, r.db.Model(&domains.User{}).Count(&count).Error
}

func (r UserRepositoryDB) SaveOrUpdateEmployee(emp *models.PSEmployee) error {
	if emp == nil {
		return fmt.Errorf("employee is nil")
	}

	domainEmp := &domains.PSEmployee{
		UHR_EmpCode:         emp.UHR_EmpCode,
		UHR_FirstName_th:    emp.UHR_FirstName_th,
		UHR_LastName_th:     emp.UHR_LastName_th,
		UHR_FullNameTh:      emp.UHR_FullNameTh,
		UHR_FirstName_en:    emp.UHR_FirstName_en,
		UHR_LastName_en:     emp.UHR_LastName_en,
		UHR_FullNameEn:      emp.UHR_FullNameEn,
		UHR_Department:      emp.UHR_Department,
		UHR_Position:        emp.UHR_Position,
		UHR_GroupDepartment: emp.UHR_GroupDepartment,
		UHR_Phone:           emp.UHR_Phone,
		UHR_OrgGroup:        emp.UHR_OrgGroup,
		UHR_OrgName:         emp.UHR_OrgName,
		AD_UserLogon:        emp.AD_UserLogon,
		AD_Mail:             emp.AD_Mail,
		AD_Phone:            emp.AD_Phone,
		AD_AccountStatus:    emp.AD_AccountStatus,
		Role:                emp.Role,
	}

	var existing domains.PSEmployee
	result := r.db.Where("UHR_EmpCode = ?", domainEmp.UHR_EmpCode).First(&existing)

	if result.Error == gorm.ErrRecordNotFound {
		return r.db.Create(domainEmp).Error
	} else if result.Error != nil {
		return result.Error
	}

	return r.db.Model(&existing).Updates(domainEmp).Error
}
