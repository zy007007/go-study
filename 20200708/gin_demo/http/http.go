package http

import (
	"gin_demo/http/routes"

	"gopkg.in/gin-gonic/gin.v1"
)

func Start() {
	router := gin.New()

	configRoutes(router)

	time.Sleep(1 * time.Second) //等待cache加载
	router.Run(config.Config().Http.Listen)
}

func configRoutes(router *gin.Engine) {
	common := router.Group("/v1/common")
	routes.ConfigCommonRoutes(common)
}

