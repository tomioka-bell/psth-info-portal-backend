package domains

import "time"

// HRSUser represents the HRS_Users table structure
type HRSUser struct {
	ID                  int        `gorm:"column:ID;primaryKey"`
	UHR_EmpCode         string     `gorm:"column:UHR_EmpCode"`
	UHR_Prefix_th       string     `gorm:"column:UHR_Prefix_th"`
	UHR_FirstName_th    string     `gorm:"column:UHR_FirstName_th"`
	UHR_LastName_th     string     `gorm:"column:UHR_LastName_th"`
	UHR_FullName_th     string     `gorm:"column:UHR_FullName_th"`
	UHR_Prefix_en       string     `gorm:"column:UHR_Prefix_en"`
	UHR_FirstName_en    string     `gorm:"column:UHR_FirstName_en"`
	UHR_LastName_en     string     `gorm:"column:UHR_LastName_en"`
	UHR_FullName_en     string     `gorm:"column:UHR_FullName_en"`
	UHR_Department      string     `gorm:"column:UHR_Department"`
	UHR_GroupDepartment string     `gorm:"column:UHR_GroupDepartment"`
	UHR_POSITION        string     `gorm:"column:UHR_POSITION"`
	UHR_IDCardCode      string     `gorm:"column:UHR_IDCardCode"`
	UHR_Province        string     `gorm:"column:UHR_Province"`
	UHR_Amphur          string     `gorm:"column:UHR_Amphur"`
	UHR_Tambon          string     `gorm:"column:UHR_Tambon"`
	UHR_Birthday        *time.Time `gorm:"column:UHR_Birthday"`
	UHR_Village         string     `gorm:"column:UHR_Village"`
	UHR_Moo             string     `gorm:"column:UHR_Moo"`
	UHR_Street          string     `gorm:"column:UHR_Street"`
	UHR_Email           string     `gorm:"column:UHR_Email"`
	UHR_PersonalEmail   string     `gorm:"column:UHR_PersonalEmail"`
	UHR_OrgCode         string     `gorm:"column:UHR_OrgCode"`
	UHR_OrgType         string     `gorm:"column:UHR_OrgType"`
	UHR_OrgName         string     `gorm:"column:UHR_OrgName"`
	UHR_OrgGroup        string     `gorm:"column:UHR_OrgGroup"`
	UHR_Shift           string     `gorm:"column:UHR_Shift"`
	UHR_ShiftDetail     string     `gorm:"column:UHR_ShiftDetail"`
	UHR_MobilePhone     string     `gorm:"column:UHR_MobilePhone"`
	UHR_Username        string     `gorm:"column:UHR_Username"`
	UHR_Password        string     `gorm:"column:UHR_Password"`
	UHR_LeaderCode      string     `gorm:"column:UHR_LeaderCode"`
	UHR_LeaderCode2     string     `gorm:"column:UHR_LeaderCode2"`
	UHR_CreateDate      *time.Time `gorm:"column:UHR_CreateDate"`
	UHR_LastDate        *time.Time `gorm:"column:UHR_LastDate"`
	UHR_Remark1         string     `gorm:"column:UHR_Remark1"`
	UHR_Remark2         string     `gorm:"column:UHR_Remark2"`
	UHR_Remark3         string     `gorm:"column:UHR_Remark3"`
	UHR_ActiveStatus    string     `gorm:"column:UHR_ActiveStatus"`
	HR_UTypeID          int        `gorm:"column:HR_UTypeID"`
	UHR_StatusToUse     string     `gorm:"column:UHR_StatusToUse"`
	HR_UTypeSection     int        `gorm:"column:HR_UTypeSection"`
	UHR_RFIDCardNo      string     `gorm:"column:UHR_RFIDCardNo"`
	UHR_RFIDRemark      string     `gorm:"column:UHR_RFIDRemark"`
	UHR_RFIDReal        string     `gorm:"column:UHR_RFIDReal"`
	UHR_RFIDRealRemark  string     `gorm:"column:UHR_RFIDRealRemark"`
	UHR_Sex             string     `gorm:"column:UHR_Sex"`
	UHR_WorkStart       *time.Time `gorm:"column:UHR_WorkStart"`
	UHR_LeaderCode3     string     `gorm:"column:UHR_LeaderCode3"`
	UHR_Company         string     `gorm:"column:UHR_Company"`
	UHR_RemarkAddUsers  string     `gorm:"column:UHR_RemarkAddUsers"`
	UHR_FileType        string     `gorm:"column:UHR_FileType"`
	UHR_UpdateBy        string     `gorm:"column:UHR_UpdateBy"`
	AD_UserLogon        string     `gorm:"column:AD_UserLogon"`
	AD_Mail             string     `gorm:"column:AD_Mail"`
	AD_Phone            string     `gorm:"column:AD_Phone"`
	AD_AccountStatus    string     `gorm:"column:AD_AccountStatus"`
}

func (HRSUser) TableName() string { return "tbl_Employees" }

// BulkEmployeeImportRequest represents bulk import request
type BulkEmployeeImportRequest struct {
	HRS_Users []HRSUser `json:"HRS_Users"`
}
