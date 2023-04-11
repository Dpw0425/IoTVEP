package app

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/db"
	"iotvep/pkg/logger"
)

func Init() *gin.Engine {
	logger.InitZap()
	config.InitConfig()
	db.InitMysql()
	db.InitRedis()
	r := InitGinEngine()
	return r
}
