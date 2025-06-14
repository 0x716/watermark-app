package routes

import (
	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	images := rg.Group(config.GlobalConfig.URL.Image.Root)
	{
		images.POST(config.GlobalConfig.URL.Image.Upload, handler.UploadImage)
		images.PUT(config.GlobalConfig.URL.Image.Apply, handler.AddWatermarks)
		images.GET(config.GlobalConfig.URL.Image.Download, handler.DownloadImages)
	}

	watermark := rg.Group(config.GlobalConfig.URL.Watermark.Root)
	{
		watermark.POST(config.GlobalConfig.URL.Watermark.Upload, handler.UploadWatermark)
		watermark.PUT(config.GlobalConfig.URL.Watermark.Apply, handler.ApplyWatermark)
		watermark.GET(config.GlobalConfig.URL.Watermark.List, handler.GetAllWatermark)
		watermark.GET(config.GlobalConfig.URL.Watermark.Get, handler.GetWatermarkById)
		watermark.DELETE(config.GlobalConfig.URL.Watermark.Delete, handler.RemoveWatermark)
	}
}
