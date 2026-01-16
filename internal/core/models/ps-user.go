package models

type UserResp struct {
	UserID    string `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname" `
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	Role      string `json:"role"`
}

type RoleInfo struct {
	Name        string `json:"role_name"`
	Description string `json:"role_description"`
}

type UserReq struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	RoleName  string `json:"role_name"`
}

type UserAdminReq struct {
	UserID    string `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	RoleName  string `json:"role_name"`
}

type UserReqAll struct {
	UserID    string `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname" `
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}

type LoginResp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginEmpResp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
