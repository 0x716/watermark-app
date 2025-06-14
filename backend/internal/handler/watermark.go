package handler

import (
	"net/http"
	"os"

	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/response"
	"github.com/0x716/watermark-app/internal/service"
	"github.com/gin-gonic/gin"
)

func UploadWatermark(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Input: "+err.Error())
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		response.Error(c, http.StatusBadRequest, "Invalid Input")
		return
	}

	watermarkService := service.NewWatermarkService()

	filenames, err := watermarkService.SaveWatermarks(files, config.GlobalConfig.Watermark.WatermarkDir, os.ModePerm)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Save watermark failed: "+err.Error())
		return
	}

	response.Success(c, filenames)
}

func ApplyWatermark(c *gin.Context) {
	var req struct {
		Id      int64   `json:"id"`
		Name    string  `json:"name"`
		Width   int64   `json:"width"`
		Height  int64   `json:"height"`
		Opacity float64 `json:"opacity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Input")
		return
	}

	watermarkService := service.NewWatermarkService()

	err := watermarkService.UpdateWatermark(req.Id, req.Name, req.Width, req.Height, float32(req.Opacity))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "update watermark info failed")
		return
	}

	response.Success(c, "update watermark success")
}

func GetWatermarkById(c *gin.Context) {
	var req struct {
		Id int64 `form:"id"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input")
		return
	}

	watermarkService := service.NewWatermarkService()

	watermark, err := watermarkService.GetWatermarkById(req.Id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "get the watermark failed")
		return
	}

	response.Success(c, watermark)
}

func GetAllWatermark(c *gin.Context) {
	watermarkService := service.NewWatermarkService()

	watermark, err := watermarkService.GetAllWatermark()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "get all watermark failed")
		return
	}

	response.Success(c, watermark)
}

func RemoveWatermark(c *gin.Context) {
	var req struct {
		Id int64 `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input")
		return
	}

	watermarkService := service.NewWatermarkService()

	err := watermarkService.DeleteWatermark(req.Id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid input")
		return
	}

	response.Success(c, "remove watermark success")
}
