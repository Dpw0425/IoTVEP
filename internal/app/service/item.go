package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/entity"
	"iotvep/internal/app/schema"
	"iotvep/pkg/error"
)

func ItemList(c *gin.Context) {
	var sil = make([]schema.ItemInfo, 0)
	result := config.MYSQLDB.Table("items").Find(&sil)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "查询失败！")
		return
	}

	if result.RowsAffected == 0 {
		error.Response(c, error.OK, gin.H{}, "暂无项目信息！")
		return
	}

	error.Response(c, error.OK, gin.H{"items": sil}, "查询成功.")
}

func AddItem(c *gin.Context, sai schema.AddItem) {
	// 判断是否存在同名项目
	var i int64
	config.MYSQLDB.Table("items").Where("item_name = ? and uid = ?", sai.ItemName, sai.UID).Count(&i)
	if i != 0 {
		error.Response(c, error.BadRequest, gin.H{}, "已存在同名项目")
		return
	}

	// 创建项目到数据库
	var ei entity.Item
	config.MYSQLDB.Table("items").Count(&i)
	ei.IID = uint(i + 1)
	ei.UID = sai.UID
	ei.ItemIMG = sai.ItemIMG
	ei.ItemName = sai.ItemName
	ei.ItemIntro = sai.ItemIntro
	result := config.MYSQLDB.Table("items").Create(&ei)
	if result.Error != nil {
		error.Response(c, error.BadRequest, gin.H{}, "新建项目失败！")
		return
	}

	// 创建项目和传感器设备关联表
	var eie entity.ItemEquipment
	for _, value := range sai.Equipments {
		eie.ID = 0
		eie.IID = uint(i + 1)
		eie.EID = uint(value)
		result1 := config.MYSQLDB.Table("item_equipments").Create(&eie)
		if result1.Error != nil {
			error.Response(c, error.BadRequest, gin.H{}, "新建项目失败！")
			config.MYSQLDB.Table("items").Where("i_id = ?", uint(i+1)).Unscoped().Delete(&ei)
			config.MYSQLDB.Table("item_equipments").Where("i_id = ?", uint(i+1)).Unscoped().Delete(&eie)
			return
		}
	}

	error.Response(c, error.OK, gin.H{}, "新建项目成功！")
}
