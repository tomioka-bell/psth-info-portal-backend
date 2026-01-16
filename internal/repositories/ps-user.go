package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
	ports "backend/internal/core/ports/repositories"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {
	if err := db.AutoMigrate(&domains.User{}); err != nil {
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
