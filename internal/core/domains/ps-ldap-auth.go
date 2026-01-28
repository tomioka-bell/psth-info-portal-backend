package domains

type PSEmployee struct {
	UHR_EmpCode         string `gorm:"column:UHR_EmpCode;primaryKey"`
	UHR_FirstName_th    string `gorm:"column:UHR_FirstName_th"`
	UHR_LastName_th     string `gorm:"column:UHR_LastName_th"`
	UHR_FullNameTh      string `gorm:"column:UHR_FullName_th"`
	UHR_FirstName_en    string `gorm:"column:UHR_FirstName_en"`
	UHR_LastName_en     string `gorm:"column:UHR_LastName_en"`
	UHR_FullNameEn      string `gorm:"column:UHR_FullName_en"`
	UHR_Department      string `gorm:"column:UHR_Department"`
	UHR_Position        string `gorm:"column:UHR_Position"`
	UHR_GroupDepartment string `gorm:"column:UHR_GroupDepartment"`
	UHR_Phone           string `gorm:"column:UHR_Phone"`
	UHR_OrgGroup        string `gorm:"column:UHR_OrgGroup"`
	UHR_OrgName         string `gorm:"column:UHR_OrgName"`
	AD_UserLogon        string `gorm:"column:AD_UserLogon;unique"`
	AD_Mail             string `gorm:"column:AD_Mail"`
	AD_Phone            string `gorm:"column:AD_Phone"`
	AD_AccountStatus    string `gorm:"column:AD_AccountStatus"`
	Role                string `gorm:"column:role;default:user"`
}

func (PSEmployee) TableName() string {
	return "ps_employees"
}
