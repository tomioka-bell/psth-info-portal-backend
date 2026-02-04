package models

type EmployeeViewResp struct {
	UHR_EmpCode         string `json:"employee_code"`
	UHR_FullNameTh      string `json:"fullname_th"`
	UHR_FullNameEn      string `json:"fullname_en"`
	UHR_FirstName_en    string `json:"firstname_en"`
	UHR_LastName_en     string `json:"lastname_en"`
	UHR_Department      string `json:"department"`
	UHR_Position        string `json:"position"`
	UHR_GroupDepartment string `json:"group_department"`
	UHR_StatusToUse     string `json:"status_to_use"`
	AD_UserLogon        string `json:"user_logon"`
	AD_Mail             string `json:"mail"`
	AD_AccountStatus    string `json:"account_status"`
	UHR_OrgGroup        string `json:"org_group"`
	UHR_OrgName         string `json:"org_name"`
	ImageURL            string `json:"image_url"`
	UHR_Phone           string `json:"phone"`
	Role                string `json:"role"`
}

type EmployeeViewByEmpCodeResp struct {
	UHR_EmpCode      string `json:"user_id"`
	UHR_FirstName_en string `json:"firstname"`
	UHR_LastName_en  string `json:"lastname"`
	UHR_Department   string `json:"role_name"`
	AD_UserLogon     string `json:"username"`
	AD_Mail          string `json:"email"`
	AD_AccountStatus string `json:"status"`
	ImageURL         string `json:"image_url"`
}

type EmployeeAdminResp struct {
	UHR_EmpCode         string `json:"employee_code"`
	UHR_FullNameTh      string `json:"fullname_th"`
	UHR_FullNameEn      string `json:"fullname_en"`
	UHR_FirstName_en    string `json:"firstname_en"`
	UHR_LastName_en     string `json:"lastname_en"`
	UHR_Department      string `json:"department"`
	UHR_Position        string `json:"position"`
	UHR_GroupDepartment string `json:"group_department"`
	UHR_StatusToUse     string `json:"status_to_use"`
	AD_UserLogon        string `json:"user_logon"`
	AD_Mail             string `json:"mail"`
	AD_AccountStatus    string `json:"account_status"`
	UHR_Phone           string `json:"phone"`
	Role                string `json:"role"`
	StatusLogin         string `json:"status_login"`
}
