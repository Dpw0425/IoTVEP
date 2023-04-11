package config

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"iotvep/pkg/logger"
)

var (
	MYSQLDB *gorm.DB
	REDISDB *redis.Client
	CONFIG  Config
)

func InitConfig() {
	viper.SetConfigFile("./config/config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("读取配置失败: %v", err)
		return
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		logger.Error("解析配置失败: %v", err)
		return
	}

	logger.Info("读取配置成功！")
}
