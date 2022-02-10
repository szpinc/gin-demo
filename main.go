package main

import (
	"gin-demo/app/routes"
	"gin-demo/bootstrap"
	"gin-demo/global"
	"github.com/gin-gonic/gin"
)

func init() {
	global.App.Log = bootstrap.InitLog()
	global.App.ConfigViper = bootstrap.InitConfig()
	global.App.DB = bootstrap.InitDB()
}

func main() {
	engine := gin.Default()
	// 初始化路由
	routes.InitRouters(engine)
	_ = engine.Run(":" + global.App.Config.App.Port)
}
