package middleware

import (
	"gin-example/pkg"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Check(name string) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		claims, err := pkg.Decode(name, token)

		logrus.WithFields(logrus.Fields{
			"claims": claims,
			"err":    err,
		}).Info("decode token")

		if err != "" {
			c.JSON(400, err)
			c.Abort()
		} else {
			data := (*claims)["data"]
			c.Set("data", data)
			c.Next()
		}
	}

}
