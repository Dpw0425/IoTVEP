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
	var sil = make([]schema.ItemIntro, 0)
	result := config.MYSQLDB.Table("items").Where("uid = ?", uid).Find(&sil)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	if result.RowsAffected == 0 {
		error.Response(c, error.OK, gin.H{}, "用户暂无已创建的项目！")
		return
	}

	error.Response(c, error.OK, gin.H{"items": sil}, "查询成功！")
}
