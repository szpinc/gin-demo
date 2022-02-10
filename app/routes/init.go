package routes

import "github.com/gin-gonic/gin"

const (
	GroupUser = "/user"
)

func InitRouters(engine *gin.Engine) {
	// 注册用户相关路由
	userRoutes(engine.Group(GroupUser))
}
