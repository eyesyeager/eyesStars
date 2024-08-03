package bootstrap

import (
	"eyesStars/app/middleware"
	"eyesStars/global"
	"eyesStars/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 17:37
 */

func setupRouter() *gin.Engine {
	if global.Native.Profiles.Active == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	// 中间件注册
	router.Use(gin.Logger(), middleware.CustomRecovery()) // 日志处理
	router.Use(middleware.Cors())                         // 跨域处理
	router.Use(middleware.Auth())                         // 权限校验

	// 设置信任代理
	// _ = router.SetTrustedProxies([]string{"192.168.1.2"})

	// 设置swagger路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册分组路由
	routes.SetContextGroupRoutes(router.Group("/context"))
	routes.SetStarGroupRoutes(router.Group("/star"))
	routes.SetTrackGroupRoutes(router.Group("/track"))
	routes.SetUserGroupRoutes(router.Group("/user"))
	routes.SetMessageGroupRoutes(router.Group("/message"))

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	_ = r.Run(":" + strconv.Itoa(global.Config.App.Port))
}
