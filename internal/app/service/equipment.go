package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/schema"
	"iotvep/pkg/error"
)

func EquipList(c *gin.Context) {
	var seil = make([]schema.EquipmentInfo, 0)
	result := config.MYSQLDB.Table("equipments").Find(&seil)

	// 本区域代码用于分页显示时控制每个页面的结果数量
	//var i int64
	//config.MYSQLDB.Table("equipments").Count(&i)

	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "获取设备列表失败！")
		return
	}

	if result.RowsAffected == 0 {
		error.Response(c, error.OK, gin.H{}, "设备列表为空！")
		return
	}

	error.Response(c, error.OK, gin.H{"equipment": seil}, "查询成功！")
}
