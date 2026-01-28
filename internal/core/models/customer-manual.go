package models

import "time"

type CreateCustomerManualRequest struct {
	CustomerManualName string `form:"customer_manual_name" binding:"required"`
	Desc               string `form:"desc" binding:"required"`
	Category           string `form:"category" binding:"required"` // office, production, quality, support
	FileName           string `form:"file_name"`
}

type UpdateCustomerManualRequest struct {
	CustomerManualName string `form:"customer_manual_name"`
	Desc               string `form:"desc"`
	Category           string `form:"category"`
	FileName           string `form:"file_name"`
}

type CustomerManualResponse struct {
	CustomerManualID   int       `json:"customer_manual_id"`
	CustomerManualName string    `json:"customer_manual_name"`
	Desc               string    `json:"desc"`
	Category           string    `json:"category"`
	FileName           string    `json:"file_name"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CustomerManualListResponse struct {
	Data       []CustomerManualResponse `json:"data"`
	Total      int64                    `json:"total"`
	Page       int                      `json:"page"`
	PageSize   int                      `json:"page_size"`
	TotalPages int                      `json:"total_pages"`
}
