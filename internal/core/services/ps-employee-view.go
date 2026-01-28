package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"backend/internal/clients"
	"backend/internal/core/models"
)

func (s *userService) SignInEmployee(dto models.LoginEmpResp) (string, error) {
	ldapResp, err := clients.LdapAuthenticate(dto.Username, dto.Password)
	if err != nil {
		return "", err
	}

	if ldapResp.Err || len(ldapResp.UserInfo) == 0 {
		return "", errors.New("authentication failed: " + ldapResp.Message)
	}

	ldapUser := ldapResp.UserInfo[0]

	if ldapUser.AD_AccountStatus != "ENABLE" {
		return "", errors.New("บัญชีนี้ถูกปิดการใช้งาน")
	}

	dbUser := models.PSEmployee{
		UHR_EmpCode:         ldapUser.EmployeeCode,
		UHR_FirstName_en:    ldapUser.FirstnameEn,
		UHR_LastName_en:     ldapUser.LastnameEn,
		UHR_FullNameTh:      ldapUser.FullnameTh,
		UHR_FullNameEn:      ldapUser.FullnameEn,
		UHR_FirstName_th:    ldapUser.FirstnameTh,
		UHR_LastName_th:     ldapUser.LastnameTh,
		UHR_Department:      ldapUser.Department,
		UHR_Position:        ldapUser.Position,
		AD_UserLogon:        ldapUser.AD_Username,
		AD_Mail:             ldapUser.AD_Mail,
		AD_AccountStatus:    ldapUser.AD_AccountStatus,
		AD_Phone:            ldapUser.AD_Phone,
		UHR_Phone:           ldapUser.AD_Phone,
		UHR_OrgGroup:        ldapUser.OrgGroup,
		UHR_OrgName:         ldapUser.OrgName,
		UHR_GroupDepartment: ldapUser.GroupDept,
		Role:                "user",
	}

	if err := s.userisrRepo.SaveOrUpdateEmployee(&dbUser); err != nil {
		return "", errors.New("เกิดข้อผิดพลาดในการเก็บข้อมูล: " + err.Error())
	}

	jwtSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	claims := jwt.MapClaims{
		"user_id":   dbUser.UHR_EmpCode,
		"username":  dbUser.AD_UserLogon,
		"firstname": dbUser.UHR_FirstName_en,
		"lastname":  dbUser.UHR_LastName_en,
		"role":      dbUser.Role,
		"status":    dbUser.AD_AccountStatus,
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

func (s *userService) GetEmployeeByEmpCodeService(userID string) (models.EmployeeViewByEmpCodeResp, error) {
	user, err := s.userisrRepo.GetEmployeeByEmpCode(userID)
	if err != nil {
		return models.EmployeeViewByEmpCodeResp{}, err
	}

	fmt.Println("user data log : ", user)

	userReq := models.EmployeeViewByEmpCodeResp{
		UHR_EmpCode:      user.UHR_EmpCode,
		UHR_FirstName_en: user.UHR_FirstName_en,
		UHR_LastName_en:  user.UHR_LastName_en,
		UHR_Department:   user.UHR_Department,
		AD_UserLogon:     user.AD_UserLogon,
		AD_Mail:          user.AD_Mail,
		AD_AccountStatus: user.AD_AccountStatus,
		ImageURL:         fmt.Sprintf("http://psth-hrservice:5020/api/v1/hrs/images/employee/%s.jpg", user.UHR_EmpCode),
	}

	return userReq, nil
}

// http://psth-hrservice:5020/api/v1/hrs/images/employee/000124.jpg
func (s *userService) GetAllEmployeesService() ([]models.EmployeeViewResp, error) {
	employees, err := s.userisrRepo.GetEmployees()
	if err != nil {
		return nil, err
	}

	var employeeResps []models.EmployeeViewResp
	for _, user := range employees {
		empResp := models.EmployeeViewResp{
			UHR_EmpCode:         user.UHR_EmpCode,
			UHR_FirstName_en:    user.UHR_FirstName_en,
			UHR_FullNameTh:      user.UHR_FullNameTh,
			UHR_FullNameEn:      user.UHR_FullNameEn,
			UHR_LastName_en:     user.UHR_LastName_en,
			UHR_Department:      user.UHR_Department,
			AD_UserLogon:        user.AD_UserLogon,
			AD_Mail:             user.AD_Mail,
			AD_AccountStatus:    user.AD_AccountStatus,
			UHR_Position:        user.UHR_Position,
			UHR_GroupDepartment: user.UHR_GroupDepartment,
			UHR_StatusToUse:     user.UHR_StatusToUse,
			UHR_Phone:           user.UHR_Phone,
			UHR_OrgGroup:        user.UHR_OrgGroup,
			UHR_OrgName:         user.UHR_OrgName,
			ImageURL:            fmt.Sprintf("http://psth-hrservice:5020/api/v1/hrs/images/employee/%s.jpg", user.UHR_EmpCode),
		}
		employeeResps = append(employeeResps, empResp)
	}

	return employeeResps, nil
}
