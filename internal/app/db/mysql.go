package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iotvep/internal/app/config"
	"iotvep/internal/app/entity"
	"iotvep/internal/app/schema"
	"iotvep/pkg/logger"
)

func InitMysql() {
	m := config.CONFIG.Mysql
	db, err := gorm.Open(mysql.Open(config.DSN(m)), &gorm.Config{})
	if err != nil {
		logger.Error("连接 Mysql 失败: %v", err)
	} else {
		// 自动建表
		db.AutoMigrate(&entity.User{})
		db.Table("equipments").AutoMigrate(&schema.EquipmentInfo{})
		db.AutoMigrate(&entity.Item{})
		db.AutoMigrate(&entity.ItemEquipment{})
		config.MYSQLDB = db
		logger.Info("连接 Mysql 成功: %v", config.DSN(m))
	}
}
