package domains

// EmployeeViewData represents the V_Employees view structure (for reading only)
type EmployeeViewData struct {
	UHR_EmpCode      string `gorm:"column:UHR_EmpCode"`
	UHR_FullName_th  string `gorm:"column:UHR_FullName_th"`
	UHR_FullName_en  string `gorm:"column:UHR_FullName_en"`
	UHR_FirstName_en string `gorm:"column:UHR_FirstName_en"`
	UHR_LastName_en  string `gorm:"column:UHR_LastName_en"`
	UHR_Department   string `gorm:"column:UHR_Department"`
	UHR_Position     string `gorm:"column:UHR_Position"`
	UHR_StatusToUse  string `gorm:"column:UHR_StatusToUse"`
	AD_UserLogon     string `gorm:"column:AD_UserLogon"`
	AD_Mail          string `gorm:"column:AD_Mail"`
	AD_Phone         string `gorm:"column:AD_Phone"`
	AD_AccountStatus string `gorm:"column:AD_AccountStatus"`
	UHR_OrgGroup     string `gorm:"column:UHR_OrgGroup"`
	UHR_OrgName      string `gorm:"column:UHR_OrgName"`
}

func (EmployeeViewData) TableName() string { return "V_Employees" }
