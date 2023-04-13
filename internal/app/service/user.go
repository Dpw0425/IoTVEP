package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/entity"
	"iotvep/internal/app/schema"
	"iotvep/pkg/error"
	"iotvep/pkg/utils"
)

func Register(c *gin.Context, su schema.User) {
	var eu entity.User
	result := config.MYSQLDB.Table("users").Where("username = ?", su.Username).First(&eu)
	if result.RowsAffected != 0 {
		error.Response(c, error.BadRequest, gin.H{}, "该用户名已存在！")
		return
	}

	if !utils.VerifyCode(su.Email, su.VerifyCode) {
		error.Response(c, error.BadRequest, gin.H{}, "验证码错误！")
		return
	}

	var eu1 entity.User
	eu1.ID = su.UID
	eu1.Username = su.Username
	eu1.Password = su.Password
	eu1.Email = su.Email
	result1 := config.MYSQLDB.Table("users").Create(&eu1)
	if result1.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "注册失败！")
		return
	}

	error.Response(c, error.OK, gin.H{}, "注册成功！")
}
