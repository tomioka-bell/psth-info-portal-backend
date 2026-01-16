package models

// BulkImportEmployeeResponse represents the response for bulk import
type BulkImportEmployeeResponse struct {
	TotalRecords   int64    `json:"total_records"`
	SuccessRecords int64    `json:"success_records"`
	FailedRecords  int64    `json:"failed_records"`
	Message        string   `json:"message"`
	Errors         []string `json:"errors,omitempty"`
}

// BulkImportEmployeeRequest represents the request for bulk import
type BulkImportEmployeeRequest struct {
	HRS_Users []map[string]interface{} `json:"HRS_Users"`
}
