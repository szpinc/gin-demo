package routes

import (
	"gin-demo/app/service/users"
	"gin-demo/pkg/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

func userRoutes(router *gin.RouterGroup) {
	router.GET("/:id", func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(200, types.SuccessResult(nil))
		}
		userInfo := users.GetUserInfo(uint(userId))
		context.JSON(200, userInfo)
	})
}
