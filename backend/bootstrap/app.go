package bootstrap

import (
	"fmt"
	"log"

	"github.com/0x716/watermark-app/internal/config"
	"github.com/0x716/watermark-app/internal/infra"
	"github.com/0x716/watermark-app/internal/router"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func init() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = infra.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("WWWWWWWWWWWWWWWWWWW")

	engine = router.InitRouter()
	infra.RegisterCleanup()
}

type Application interface {
	Run() error
}

type ApplicationImpl struct {
	Engine *gin.Engine
}

func NewApplication() Application {
	return &ApplicationImpl{Engine: engine}
}

func (a *ApplicationImpl) Run() error {
	if err := a.Engine.Run(fmt.Sprintf("%s:%s", config.GlobalConfig.App.Host, config.GlobalConfig.App.Port)); err != nil {
		return err
	}

	return nil
}
