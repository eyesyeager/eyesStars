package routes

import (
	"eyesStars/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2024/7/25 22:39
 */

func SetMessageGroupRoutes(router *gin.RouterGroup) {
	// 发送回忆邮件
	router.POST("/sendMemoryEmail", controller.SendMemoryEmail)
}
