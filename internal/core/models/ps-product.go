package models

import "time"

type CreateProductRequest struct {
	ProductName   string `json:"product_name"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	ProductMainImages string `json:"product_main_images"`
	ProductImages string `json:"product_images"`
	Recommend   bool   `json:"recommend"`
}

type UpdateProductRequest struct {
	ProductName       string   `json:"product_name"`
	Category          string   `json:"category"`
	Description       string   `json:"description"`
	ProductMainImages string   `json:"product_main_images"`
	ProductImages     string   `json:"product_images"`
	ImagesToDelete    []string `json:"images_to_delete"`
	DeleteMainImage   bool     `json:"delete_main_image"`
	Recommend         *bool    `json:"recommend"`
}

type CreateProductResponse struct {
	ProductID     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	ProductMainImages string `json:"product_main_images"`
	ProductImages string `json:"product_images"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Recommend   bool      `json:"recommend"`
}

type ProductResponse struct {
	ProductID     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	ProductMainImages string `json:"product_main_images"`
	ProductImages string `json:"product_images"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Recommend   bool      `json:"recommend"`
}