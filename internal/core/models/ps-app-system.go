package models

import "time"

type AppSystemCategoryRequest struct {
	Category string                   `json:"category"`
	Systems  []CreateAppSystemRequest `json:"systems"`
}

type CreateAppSystemRequest struct {
	Name      string                   `json:"name" binding:"required"`
	Desc      string                   `json:"desc" binding:"required"`
	Category  string                   `json:"category" binding:"required"` // office, production, quality, support, hr
	Href      string                   `json:"href" binding:"required"`
	Icon      string                   `json:"icon"` // React icon name or base64 encoded icon
	ParentID  *int                     `json:"parent_id"`
	SortOrder int                      `json:"sort_order"`
	Children  []CreateAppSystemRequest `json:"children,omitempty"` // nested children สำหรับ recursive create
}

// CreateAppSystemNestedRequest สำหรับรับ nested structure
type CreateAppSystemNestedRequest struct {
	Name      string                         `json:"name" binding:"required"`
	Desc      string                         `json:"desc" binding:"required"`
	Category  string                         `json:"category" binding:"required"`
	Href      string                         `json:"href" binding:"required"`
	Icon      string                         `json:"icon"`
	SortOrder int                            `json:"sort_order"`
	Children  []CreateAppSystemNestedRequest `json:"children,omitempty"`
}

type UpdateAppSystemRequest struct {
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Category      string `json:"category"`
	Href          string `json:"href"`
	Icon          string `json:"icon"`
	ParentID      *int   `json:"parent_id"`
	SortOrder     int    `json:"sort_order"`
	ClearParentID bool   `json:"clear_parent_id"` // Flag to indicate parent_id should be set to NULL
}

type AppSystemResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Category  string    `json:"category"`
	Href      string    `json:"href"`
	Icon      string    `json:"icon"`
	ParentID  int       `json:"parent_id"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AppSystemListResponse struct {
	Data       []AppSystemResponse `json:"data"`
	Total      int64               `json:"total"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
	TotalPages int                 `json:"total_pages"`
}
