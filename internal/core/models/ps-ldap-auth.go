package models

type LdapUserInfo struct {
	EmployeeCode     string  `json:"employee_code"`
	PrefixTh         string  `json:"prefix_th"`
	FirstnameTh      string  `json:"firstname_th"`
	LastnameTh       string  `json:"lastname_th"`
	FullnameTh       string  `json:"fullname_th"`
	PrefixEn         string  `json:"prefix_en"`
	FirstnameEn      string  `json:"firstname_en"`
	LastnameEn       string  `json:"lastname_en"`
	FullnameEn       string  `json:"fullname_en"`
	Sex              string  `json:"sex"`
	Department       string  `json:"department"`
	Position         string  `json:"position"`
	AD_Username      string  `json:"ad_username"`
	AD_Mail          string  `json:"ad_mail"`
	AD_Phone         string  `json:"ad_phone"`
	AD_AccountStatus string  `json:"ad_account_status"`
	WorkStart        string  `json:"work_start"`
	WorkEnd          *string `json:"work_end"`
	OrgCode          string  `json:"org_code"`
	GroupDept        string  `json:"group_dept"`
	OrgGroup         string  `json:"org_group"`
	OrgName          string  `json:"org_name"`
	PersonalMail     string  `json:"personal_mail"`
}

type LdapAuthResponse struct {
	Err      bool           `json:"err"`
	Message  string         `json:"message"`
	UserInfo []LdapUserInfo `json:"user_info"`
}

type PSEmployee struct {
	UHR_EmpCode         string `json:"UHR_EmpCode"`
	UHR_FirstName_th    string `json:"UHR_FirstName_th"`
	UHR_LastName_th     string `json:"UHR_LastName_th"`
	UHR_FullNameTh      string `json:"UHR_FullNameTh"`
	UHR_FirstName_en    string `json:"UHR_FirstName_en"`
	UHR_LastName_en     string `json:"UHR_LastName_en"`
	UHR_FullNameEn      string `json:"UHR_FullNameEn"`
	UHR_Department      string `json:"UHR_Department"`
	UHR_Position        string `json:"UHR_Position"`
	UHR_GroupDepartment string `json:"UHR_GroupDepartment"`
	UHR_Phone           string `json:"UHR_Phone"`
	UHR_OrgGroup        string `json:"UHR_OrgGroup"`
	UHR_OrgName         string `json:"UHR_OrgName"`
	AD_UserLogon        string `json:"AD_UserLogon"`
	AD_Mail             string `json:"AD_Mail"`
	AD_Phone            string `json:"AD_Phone"`
	AD_AccountStatus    string `json:"AD_AccountStatus"`
	Role                string `json:"Role"`
}
