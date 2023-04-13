package api

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/service"
)

func SendEmailCode(c *gin.Context) {
	email := c.Query("email")

	service.SendEmailCode(c, email)
}
