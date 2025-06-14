package handler

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/response"
	"github.com/0x716/watermark-app/internal/service"
	"github.com/0x716/watermark-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invaild form data")
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		response.Error(c, http.StatusBadRequest, "Please at least upload one image")
		return
	}

	filenames, err := service.SaveUploadImages(files, config.GlobalConfig.Image.ImageDir, os.ModePerm)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Image proccess failed: "+err.Error())
		return
	}

	response.Success(c, filenames)
}

func AddWatermarks(c *gin.Context) {
	var req struct {
		FileNames     []string `json:"fileNames"`
		WatermarkName string   `json:"watermakrName"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	if len(req.FileNames) == 0 || req.WatermarkName == "" {
		response.Error(c, http.StatusBadRequest, "invalid iput")
	}

	outputPaths := make([]string, 0, config.GlobalConfig.Constants.StringLength)

	for _, fileName := range req.FileNames {
		imagePath := filepath.Join(config.GlobalConfig.Image.ImageDir, fileName)
		watermarkPath := filepath.Join(config.GlobalConfig.Watermark.WatermarkDir, req.WatermarkName)
		outputPath := filepath.Join(config.GlobalConfig.Image.OutputDir, utils.GenerateFilename(fileName))

		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			response.Error(c, http.StatusNotFound, "Image not found")
			return
		}

		if _, err := os.Stat(watermarkPath); os.IsNotExist(err) {
			response.Error(c, http.StatusNotFound, "Watermark not found")
			return
		}

		if err := service.AddWatermark(imagePath, watermarkPath, outputPath, config.GlobalConfig.Watermark.Opacity, config.GlobalConfig.Watermark.Scale); err != nil {
			response.Error(c, http.StatusInternalServerError, "Watermark failed")
			return
		}

		outputPaths = append(outputPaths, outputPath)
	}

	response.Success(c, outputPaths)
}

func DownloadImages(c *gin.Context) {
	var req struct {
		Files []string `json:"files"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Input: "+err.Error())
		return
	}

	if len(req.Files) == 0 {
		response.Error(c, http.StatusBadRequest, "Invalid Input")
		return
	}

	// Set HTTP Header
	c.Writer.Header().Set("Content-Type", "application/zip")                                    // Set Response Content Type As Zip File
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=watermarked_images.zip") // Tell Browser This Is a Attachment, Please Download It, And Not Open It.
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")                                // This is binary transfer.
	c.Writer.Header().Set("Cache-Control", "no-cache")                                          // Don't Cache This Response

	err := service.CreateZipFromFiles(req.Files, config.GlobalConfig.Image.OutputDir, c.Writer)
	if err != nil {
		log.Printf("failed to create zip: %v", err)
		return
	}

	response.Success(c, "download images success")
}
