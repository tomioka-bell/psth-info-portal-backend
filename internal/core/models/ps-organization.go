package models

import "time"

type CreateOrganizationRequest struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc" binding:"required"`
	Category string `json:"category" binding:"required"` // office, production, quality, support
	Href     string `json:"href" binding:"required"`
	Icon     string `json:"icon"` // React icon name or base64 encoded icon
}

type UpdateOrganizationRequest struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Category string `json:"category"`
	Href     string `json:"href"`
	Icon     string `json:"icon"`
	FileName string `json:"file_name"`
}

type OrganizationResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Category  string    `json:"category"`
	Href      string    `json:"href"`
	Icon      string    `json:"icon"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrganizationListResponse struct {
	Data       []OrganizationResponse `json:"data"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
	TotalPages int                    `json:"total_pages"`
}
