package routes

import (
	"gin_demo/controller"
	"gopkg.in/gin-gonic/gin.v1"
)

//通用的路由
func ConfigCommonRoutes(r *gin.RouterGroup) {
	r.GET("/test", controller.CommonControllerRepo.Test)
}
