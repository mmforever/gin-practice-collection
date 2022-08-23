package main

import (
	"gin-example/model"
	"gin-example/pkg"
	"gin-example/router"
)

func main() {
	pkg.LogSetup()
	pkg.AppSetup()
	model.Setup()
	pkg.RedisSetup()
	pkg.JwtSetup()
	r := router.InitRouter()

	r.Run(":" + pkg.AppConfig.Port)
}
