package domains

type EmployeeView struct {
	UHR_EmpCode         string `gorm:"column:UHR_EmpCode"`
	UHR_FullNameTh      string `gorm:"column:UHR_FullName_th"`
	UHR_FullNameEn      string `gorm:"column:UHR_FullName_en"`
	UHR_FirstName_en    string `gorm:"column:UHR_FirstName_en"`
	UHR_LastName_en     string `gorm:"column:UHR_LastName_en"`
	UHR_Department      string `gorm:"column:UHR_Department"`
	UHR_Position        string `gorm:"column:UHR_POSITION"`
	UHR_Phone           string `gorm:"column:UHR_Phone"`
	UHR_GroupDepartment string `gorm:"column:UHR_GroupDepartment"`
	UHR_StatusToUse     string `gorm:"column:UHR_StatusToUse"`
	AD_UserLogon        string `gorm:"column:AD_UserLogon"`
	AD_Mail             string `gorm:"column:UHR_Email"`
	AD_Phone            string `gorm:"column:AD_Phone"`
	AD_AccountStatus    string `gorm:"column:AD_AccountStatus"`
	UHR_OrgGroup        string `gorm:"column:UHR_OrgGroup"`
	UHR_OrgName         string `gorm:"column:UHR_OrgName"`

	TempOTP string `gorm:"-"`
}

func (EmployeeView) TableName() string { return "PSTH_HR_SERVICE.dbo.V_HRS_UsersHrService" }
