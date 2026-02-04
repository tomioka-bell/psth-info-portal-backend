package models

import "time"

type CreateSafetyDocumentRequest struct {
	SafetyDocumentName string `json:"safety_document_name"`
	SafetyDocumentDesc string `json:"safety_document_desc"`
	Category           string `json:"category"`
	Department         string `json:"department"`
	FileName           string `json:"file_name"`
}

type UpdateSafetyDocumentRequest struct {
	SafetyDocumentName string `json:"safety_document_name"`
	SafetyDocumentDesc string `json:"safety_document_desc"`
	Category           string `json:"category"`
	Department         string `json:"department"`
	FileName           string `json:"file_name"`
}

type SafetyDocumentResponse struct {
	SafetyDocumentID   int       `json:"safety_document_id"`
	SafetyDocumentName string    `json:"safety_document_name"`
	SafetyDocumentDesc string    `json:"safety_document_desc"`
	Category           string    `json:"category"`
	Department         string    `json:"department"`
	FileName           string    `json:"file_name"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type SafetyDocumentListResponse struct {
	Data       []SafetyDocumentResponse `json:"data"`
	Total      int64                    `json:"total"`
	Page       int                      `json:"page"`
	PageSize   int                      `json:"page_size"`
	TotalPages int                      `json:"total_pages"`
}
