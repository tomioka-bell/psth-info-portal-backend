package models

import "time"

type CreateCustomerManualRequest struct {
	CustomerManualName string                        `form:"customer_manual_name" json:"customer_manual_name" binding:"required"`
	Desc               string                        `form:"desc" json:"desc" binding:"required"`
	Category           string                        `form:"category" json:"category" binding:"required"` // office, production, quality, support
	FileName           string                        `form:"file_name" json:"file_name"`
	ParentID           *int                          `form:"parent_id" json:"parent_id"`
	SortOrder          int                           `form:"sort_order" json:"sort_order"`
	Children           []CreateCustomerManualRequest `form:"children" json:"children"`
}

type UpdateCustomerManualRequest struct {
	CustomerManualName string `form:"customer_manual_name" json:"customer_manual_name"`
	Desc               string `form:"desc" json:"desc"`
	Category           string `form:"category" json:"category"`
	FileName           string `form:"file_name" json:"file_name"`
	ParentID           *int   `form:"parent_id" json:"parent_id"`
	SortOrder          int    `form:"sort_order" json:"sort_order"`
	ClearParentID      bool   `form:"clear_parent_id" json:"clear_parent_id"` // Flag to indicate parent_id should be set to NULL
}

type CustomerManualResponse struct {
	CustomerManualID   int       `json:"customer_manual_id"`
	CustomerManualName string    `json:"customer_manual_name"`
	Desc               string    `json:"desc"`
	Category           string    `json:"category"`
	FileName           string    `json:"file_name"`
	ParentID           *int      `json:"parent_id"`
	SortOrder          int       `json:"sort_order"`
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
