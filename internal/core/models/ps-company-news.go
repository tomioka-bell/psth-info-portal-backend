package models

import "github.com/google/uuid"

type CompanyNewsReq struct {
	CompanyNewsID    uuid.UUID `json:"company_news_id"`
	CompanyNewsPhoto string    `json:"company_news_photo"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	Category         string    `json:"category"`
	UsernameCreator  string    `json:"username_creator"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
}

type CompanyNewsListResp struct {
	Total      int64            `json:"total"`
	TotalPages int              `json:"total_pages"`
	Limit      int              `json:"limit"`
	Offset     int              `json:"offset"`
	Data       []CompanyNewsReq `json:"data"`
}

type CompanyNewsResp struct {
	CompanyNewsID    string `json:"company_news_id"`
	CompanyNewsPhoto string `json:"company_news_photo"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Category         string `json:"category"`
	UsernameCreator  string `json:"username_creator"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
