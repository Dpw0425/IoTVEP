package router

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/middleware"
)

func Register(app *gin.Engine) error {
	RegisterAPI(app)
	return nil
}

func RegisterAPI(app *gin.Engine) {
	app.Use(middleware.LimitRoute(config.CONFIG.Limit.Limit))
	{

	}
}
