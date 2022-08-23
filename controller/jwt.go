package controller

import (
	"gin-example/pkg"

	"github.com/gin-gonic/gin"
)

func Parse(c *gin.Context) {
	data, ex := c.Get("data")

	if !ex {
		c.JSON(400, "服务器错误")
	}
	c.JSON(200, data)
}

func Token(c *gin.Context) {
	pkg.JwtSetup()
	name := "merchant"
	data := map[string]string{"name": c.Query("username")}

	ret := make(map[string]string)
	ret["token"] = pkg.Encode(name, data)
	c.JSON(200, ret)

}
