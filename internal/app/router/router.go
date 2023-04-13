package router

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/api"
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
		user := app.Group("/user")
		{
			user.GET("/verify", api.SendEmailCode)
			user.POST("/register", api.Register)
			user.POST("/login", api.Login)
		}
	}
}
