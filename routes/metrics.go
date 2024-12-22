package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:52
 */

func SetMetricsGroupRoutes(router *gin.RouterGroup) {
	router.GET("/prometheus", gin.WrapH(promhttp.Handler()))
}
