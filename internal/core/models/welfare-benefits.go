package models

import "time"

type CreateWelfareBenefitRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	FileName    string `json:"file_name"`
	ImageURL    string `json:"image_url"`
}

type UpdateWelfareBenefitRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	FileName    string `json:"file_name"`
	ImageURL    string `json:"image_url"`
}

type WelfareBenefitResponse struct {
	WelfareBenefitID int       `json:"welfare_benefit_id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Category         string    `json:"category"`
	ImageURL         string    `json:"image_url"`
	FileName         string    `json:"file_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type WelfareBenefitListResponse struct {
	Data       []WelfareBenefitResponse `json:"data"`
	Total      int64                    `json:"total"`
	Page       int                      `json:"page"`
	PageSize   int                      `json:"page_size"`
	TotalPages int                      `json:"total_pages"`
}
