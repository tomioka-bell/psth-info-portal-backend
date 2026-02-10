package models

import "time"

type CreateQmsDocumentRequest struct {
	QmsDocumentsName  string `form:"qms_documents_name" binding:"required"`
	DQmsDocumentsDesc string `form:"dqms_documents_desc" binding:"required"`
	Category          string `form:"category" binding:"required"`
	FileName          string `form:"file_name"`
}

type UpdateQmsDocumentRequest struct {
	QmsDocumentsName  string `form:"qms_documents_name"`
	DQmsDocumentsDesc string `form:"dqms_documents_desc"`
	Category          string `form:"category"`
	FileName          string `form:"file_name"`
}

type QmsDocumentResponse struct {
	QmsDocumentsID    int       `json:"qms_documents_id"`
	QmsDocumentsName  string    `json:"qms_documents_name"`
	DQmsDocumentsDesc string    `json:"dqms_documents_desc"`
	Category          string    `json:"category"`
	FileName          string    `json:"file_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type QmsDocumentListResponse struct {
	Data       []QmsDocumentResponse `json:"data"`
	Total      int64                 `json:"total"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
	TotalPages int                   `json:"total_pages"`
}
