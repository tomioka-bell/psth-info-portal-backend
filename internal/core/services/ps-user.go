package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	ports "backend/internal/core/ports/repositories"
	servicesports "backend/internal/core/ports/services"
	"backend/internal/pkgs/logs"
	"backend/internal/pkgs/utils"
)

type userService struct {
	userisrRepo ports.UserRepository
}

func NewUserService(UserisrRepo ports.UserRepository) servicesports.UserService {
	return &userService{userisrRepo: UserisrRepo}
}

func (s *userService) CreateUserService(req models.UserResp) error {

	newID := uuid.New()

	encodedPassword := utils.Encode(req.Password)

	domainISR := domains.User{
		UserID:    newID.String(),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Username:  req.Username,
		Email:     req.Email,
		Password:  encodedPassword,
		Status:    "activate",
	}

	if s.userisrRepo == nil {
		log.Println("UserRepo is nil")
		return fmt.Errorf("user repository is not initialized")
	}

	err := s.userisrRepo.CreateUserRepository(&domainISR)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *userService) UpdateUserWithMapService(userID string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	if password, ok := updates["password"]; ok {
		encodedPassword := utils.Encode(password.(string))
		updates["password"] = encodedPassword
	}

	return s.userisrRepo.UpdateUserWithMap(userID, updates)
}

func (s *userService) SignIn(dto models.LoginEmpResp) (string, error) {
	user, err := s.userisrRepo.FindByUsername(dto.Username)
	if err != nil {
		return "", errors.New("ชื่อผู้ใช้ไม่ถูกต้อง")
	}

	if user == nil {
		return "", errors.New("ไม่พบผู้ใช้")
	}

	if user.Status == "disable" {
		return "", errors.New("บัญชีนี้ถูกปิดการใช้งาน")
	}

	jwtSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	claims := jwt.MapClaims{
		"user_id":   user.UserID,
		"username":  user.Username,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
		"status":    user.Status,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", errors.New("เกิดข้อผิดพลาดในการเซ็นชื่อ JWT")
	}

	return signedToken, nil
}

// func (s *userService) SignIn(dto models.LoginEmpResp) (string, error) {
// 	user, err := s.userisrRepo.FindByUsername(dto.Username)
// 	if err != nil {
// 		return "", errors.New("ชื่อผู้ใช้ไม่ถูกต้อง")
// 	}

// 	if user == nil {
// 		return "", errors.New("ไม่พบผู้ใช้")
// 	}

// 	if user.Status == "disable" {
// 		return "", errors.New("บัญชีนี้ถูกปิดการใช้งาน")
// 	}

// 	if !utils.Compare(dto.Password, user.Password) {
// 		return "", errors.New("รหัสผ่านไม่ถูกต้อง")
// 	}

// 	jwtSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
// 	claims := jwt.MapClaims{
// 		"user_id":   user.UserID,
// 		"username":  user.Username,
// 		"firstname": user.Firstname,
// 		"lastname":  user.Lastname,
// 		"status":    user.Status,
// 		"iat":       time.Now().Unix(),
// 		"exp":       time.Now().Add(time.Hour * 24).Unix(),
// 	}

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := jwtToken.SignedString(jwtSecretKey)
// 	if err != nil {
// 		return "", errors.New("เกิดข้อผิดพลาดในการเซ็นชื่อ JWT")
// 	}

// 	return signedToken, nil
// }

func (s *userService) GetProfileByCookieId(userID string) (models.UserReq, error) {
	// ดึงข้อมูล User จาก repository
	user, err := s.userisrRepo.GetUserByID(userID)
	if err != nil {
		return models.UserReq{}, err
	}

	// สร้างโครงสร้าง UserReq
	userReq := models.UserReq{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		RoleName:  user.Role,
	}

	// เตรียมข้อมูล RoleName
	// var roleInfoList []models.RoleInfo
	// for _, role := range user.Role {
	// 	roleInfo := models.RoleInfo{
	// 		Name:        role.RoleName,
	// 		Description: role.RoleDescription,
	// 	}
	// 	roleInfoList = append(roleInfoList, roleInfo)
	// }
	// userReq.RoleName = roleInfoList

	return userReq, nil
}

func (s *userService) GetUserByID(userID string) (models.UserAdminReq, error) {
	// ดึงข้อมูล User จาก repository
	user, err := s.userisrRepo.GetUserByID(userID)
	if err != nil {
		return models.UserAdminReq{}, err
	}

	// สร้างโครงสร้าง UserReq
	userReq := models.UserAdminReq{
		UserID:    user.UserID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		RoleName:  user.Role,
	}

	// เตรียมข้อมูล RoleName
	// var roleInfoList []models.RoleInfo
	// for _, role := range user.Role {
	// 	roleInfo := models.RoleInfo{
	// 		Name: role.RoleName,
	// 	}
	// 	roleInfoList = append(roleInfoList, roleInfo)
	// }
	// userReq.RoleName = roleInfoList

	return userReq, nil
}

func (s *userService) GetAllUserSevice() ([]models.UserReqAll, error) {
	domainRequests, err := s.userisrRepo.GetAllUser()
	if err != nil {
		logs.Error(err)
		return nil, errors.New("failed to get all user")
	}

	var requests []models.UserReqAll

	for _, domainRequest := range domainRequests {
		// if domainRequest.Status == "Disable" {
		// 	continue
		// }

		request := models.UserReqAll{
			UserID:    domainRequest.UserID,
			Firstname: domainRequest.Firstname,
			Lastname:  domainRequest.Lastname,
			Username:  domainRequest.Username,
			Email:     domainRequest.Email,
			Status:    domainRequest.Status,
		}
		requests = append(requests, request)
	}

	return requests, nil
}

func (s *userService) GetUserCountService() (int64, error) {
	count, err := s.userisrRepo.GetUserCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}
