package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/pkg/error"
	"iotvep/pkg/utils"
)

func SendEmailCode(c *gin.Context, email string) {
	if utils.GenerateCode(email) {
		error.Response(c, error.OK, gin.H{}, "验证码发送成功.")
	} else {
		error.Response(c, error.BadRequest, gin.H{}, "验证码发送失败.")
	}
}
