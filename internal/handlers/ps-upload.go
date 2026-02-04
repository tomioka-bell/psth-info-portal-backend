package handlers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func ServeUploadFile(c *fiber.Ctx) error {

	folder := c.Query("folder")
	filename := c.Query("filename")

	if folder == "" || filename == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid path")
	}
	filePath := filepath.Join("uploads", folder, filename)

	log.Printf("[ServeUploadFile] Attempting to serve file: %s\n", filePath)

	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("[ServeUploadFile] File not found: %s\n", filePath)
		dirPath := filepath.Join("uploads", folder)
		files, _ := os.ReadDir(dirPath)
		log.Printf("[ServeUploadFile] Files in %s:\n", dirPath)
		for _, f := range files {
			log.Printf("  - %s\n", f.Name())
		}
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("File not found: %s", filePath))
	}

	if err != nil {
		log.Printf("[ServeUploadFile] Error checking file: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error accessing file: %v", err))
	}

	log.Printf("[ServeUploadFile] File found, size: %d bytes\n", fileInfo.Size())

	return c.SendFile(filePath)
}
