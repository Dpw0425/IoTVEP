package app

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/middleware"
	"iotvep/internal/app/router"
)

func InitGinEngine() *gin.Engine {
	gin.SetMode(config.CONFIG.Mode.RunMode)

	app := gin.Default()

	// CORS
	if config.CONFIG.Cors.Enable {
		app.Use(middleware.Cors())
	}

	// Router register
	router.Register(app)

	return app
}
