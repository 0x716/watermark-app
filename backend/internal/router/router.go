package router

import (
	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/router/routes"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	engine.Static(config.GlobalConfig.URL.Output, config.GlobalConfig.Image.OutputDir)

	apiGroup := engine.Group(config.GlobalConfig.URL.Root)
	routes.RegisterRoutes(apiGroup)

	return engine
}
