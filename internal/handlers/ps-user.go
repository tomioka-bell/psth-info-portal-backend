package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	"backend/internal/clients"
	"backend/internal/core/domains"
	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
)

type UserHandler struct {
	UserSrv services.UserService
}

func NewUserHandler(insSrv services.UserService) *UserHandler {
	return &UserHandler{UserSrv: insSrv}
}

func (h *UserHandler) LoginDBHandler(c *fiber.Ctx) error {
	var loginData models.LoginEmpResp
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
	}

	token, err := h.UserSrv.SignIn(loginData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "PPR_",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		// SameSite: "Lax",
		Secure:   false,
		SameSite: "none",
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Login": "เข้าสู่ระบบสำเร็จ",
	})
}

func (h *UserHandler) LoginHandler(c *fiber.Ctx) error {
	var loginData models.LoginEmpResp
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
	}

	// 1) LDAP auth
	ok, msg := clients.LdapAuthenticate(loginData.Username, loginData.Password)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	// token, err := h.UserSrv.SignIn(loginData)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	token, err := h.UserSrv.SignInEmployee(loginData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "เข้าสู่ระบบสำเร็จ",
		"token":   token,
	})
}

func (h *UserHandler) LogoutDBHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "PPR_",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "none",
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Logout": "ออกจากระบบสำเร็จ",
	})
}

func (h *UserHandler) GetProfileCookie(c *fiber.Ctx) error {
	cookie := c.Cookies("ACG_")

	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(viper.GetString("token.secret_key")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)

		result, err := h.UserSrv.GetProfileByCookieId(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"result": result,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized access",
	})
}

func (h *UserHandler) GetProfileHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	tokenString := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization header",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(viper.GetString("token.secret_key")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	userID := claims["user_id"].(string)

	result, err := h.UserSrv.GetEmployeeByEmpCodeService(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func (h *UserHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "กรุณาระบุ user_id"})
	}

	companyContents, err := h.UserSrv.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(companyContents)
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var req models.UserResp

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if h.UserSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	err := h.UserSrv.CreateUserService(req)
	if err != nil {
		log.Println("Error creating  User:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create  User",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User  created successfully",
	})
}

func (h *UserHandler) GetAllUserHandler(c *fiber.Ctx) error {
	Companies, err := h.UserSrv.GetAllUserSevice()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to user ",
		})
	}

	return c.Status(fiber.StatusOK).JSON(Companies)
}

func (h *UserHandler) GetAllEmployeesHandler(c *fiber.Ctx) error {
	employees, err := h.UserSrv.GetAllEmployeesService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve employees",
		})
	}

	return c.Status(fiber.StatusOK).JSON(employees)
}

func (h *UserHandler) CheckAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("Car_")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(viper.GetString("token.secret_key")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"]
		c.Locals("user_id", userID)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized access",
	})
}

func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	var updatedQuestion domains.User
	if err := c.BodyParser(&updatedQuestion); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	updates := map[string]interface{}{}

	if updatedQuestion.Firstname != "" {
		updates["firstname"] = updatedQuestion.Firstname
	}
	if updatedQuestion.Lastname != "" {
		updates["lastname"] = updatedQuestion.Lastname
	}
	if updatedQuestion.Username != "" {
		updates["username"] = updatedQuestion.Username
	}
	if updatedQuestion.Email != "" {
		updates["email"] = updatedQuestion.Email
	}
	if updatedQuestion.Password != "" {
		updates["password"] = updatedQuestion.Password
	}
	if updatedQuestion.Status != "" {
		updates["status"] = updatedQuestion.Status
	}

	err := h.UserSrv.UpdateUserWithMapService(userID, updates)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *UserHandler) GetUserCountHandler(c *fiber.Ctx) error {
	count, err := h.UserSrv.GetUserCountService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user count",
		})
	}

	return c.JSON(fiber.Map{
		"total_users": count,
	})
}
