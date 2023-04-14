package api

import (
	"github.com/gin-gonic/gin"
	"iotvep/internal/app/service"
	"strconv"
)

func Statistics(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))

	service.Statistics(c, uint(uid))
}

func MyItem(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))

	service.MyItem(c, uint(uid))
}
