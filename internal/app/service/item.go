package service

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/config"
	"iotvep/internal/app/schema"
	"iotvep/pkg/error"
)

func ItemList(c *gin.Context) {
	var sil = make([]schema.ItemIntro, 0)
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
