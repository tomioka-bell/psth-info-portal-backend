package models

import "time"

type CreateOrganizationDocRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Department string `json:"department"`
	FileName   string `json:"file_name"`
}

type UpdateOrganizationDocRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Department string `json:"department"`
	FileName   string `json:"file_name"`
}

type OrganizationDocResponse struct {
	OrganizationDocID int       `json:"organization_doc_id"`
	Name              string    `json:"name"`
	Desc              string    `json:"desc"`
	Department        string    `json:"department"`
	FileName          string    `json:"file_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type OrganizationDocListResponse struct {
	Data       []OrganizationDocResponse `json:"data"`
	Total      int64                     `json:"total"`
	Page       int                       `json:"page"`
	PageSize   int                       `json:"page_size"`
	TotalPages int                       `json:"total_pages"`
}
