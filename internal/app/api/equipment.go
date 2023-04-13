package api

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/service"
)

// 本项目在此阶段暂时无需使用该函数传输数据
// 保留此函数是方便后续修改
func EquipList(c *gin.Context) {
	service.EquipList(c)
}
