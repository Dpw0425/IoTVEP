package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/schema"
	"iotvep/pkg/error"
)

func Statistics(c *gin.Context, uid uint) {
	var ss schema.Statistics
	config.MYSQLDB.Table("users").Count(&ss.UserCount)
	config.MYSQLDB.Table("equipments").Count(&ss.EquipmentCount)
	config.MYSQLDB.Table("items").Where("uid = ?", uid).Count(&ss.ItemCount)

	error.Response(c, error.OK, gin.H{"statistics": ss}, "查询成功！")
}

func MyItem(c *gin.Context, uid uint) {

}
