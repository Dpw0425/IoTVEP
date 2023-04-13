package api

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/schema"
	"iotvep/internal/app/service"
	"iotvep/pkg/error"
)

func Register(c *gin.Context) {
	var su schema.User
	if err := c.ShouldBindJSON(&su); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.Register(c, su)
}

func Login(c *gin.Context) {
	var su schema.User
	if err := c.ShouldBindJSON(&su); err != nil {
		error.Response(c, error.BadRequest, gin.H{}, "输入内容有误！")
		return
	}

	service.Login(c, su)
}
