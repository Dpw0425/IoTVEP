package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/entity"
	"iotvep/internal/app/schema"
	"iotvep/pkg/error"
	"iotvep/pkg/logger"
	"iotvep/pkg/utils"
)

func Register(c *gin.Context, su schema.User) {
	// 判断用户名是否存在
	var eu entity.User
	result := config.MYSQLDB.Table("users").Where("username = ?", su.Username).First(&eu)
	if result.RowsAffected != 0 {
		error.Response(c, error.BadRequest, gin.H{}, "该用户名已存在！")
		return
	}

	// 判断验证码是否正确
	if !utils.VerifyCode(su.Email, su.VerifyCode) {
		error.Response(c, error.BadRequest, gin.H{}, "验证码错误！")
		return
	}

	// 保存到数据库
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

func Login(c *gin.Context, su schema.User) {
	// 判断用户名是否存在
	var eu entity.User
	result := config.MYSQLDB.Table("users").Where("username = ?", su.Username).First(&eu)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "用户尚未注册！")
		return
	}

	// 验证密码是否正确
	if !(su.Password == eu.Password) {
		error.Response(c, error.BadRequest, gin.H{}, "密码错误！")
		return
	}

	// 分发 token
	token, err := utils.ReleaseToken(eu)
	if err != nil {
		logger.Error("Token 发送失败！", err)
		error.Response(c, error.BadRequest, gin.H{}, "token 发送失败！")
		return
	}

	error.Response(c, error.OK, gin.H{"token": token}, "登录成功！")
}
