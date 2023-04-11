package main

import (
	"iotvep/internal/app"
	"iotvep/internal/app/config"
)

func main() {
	r := app.Init()
	r.Run(":" + config.CONFIG.Http.Port) // 启动服务
}
