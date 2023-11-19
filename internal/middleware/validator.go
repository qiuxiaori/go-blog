package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {
		param, exists := c.Get("Payload.Param")
		if exists {
			valid, _ := app.BindAndValid(c, &param)
			if !valid {
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
