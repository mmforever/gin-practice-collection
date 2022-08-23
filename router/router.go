package router

import (
	"gin-example/middleware"

	"gin-example/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	gin.SetMode("debug")

	r.POST("/parse", middleware.Check("merchant"), controller.Parse)

	r.GET("/token", controller.Token)

	return r

}
