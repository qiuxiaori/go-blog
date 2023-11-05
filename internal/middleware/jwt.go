package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var (
		// 	token string
		// 	ecode = errcode.Success
		// )
		res, err := app.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBLZXkiOiJhMDE2MTAyMjhmZTk5OGY1MTVhNzJkZDczMDI5NGQ4NyIsIkFwcFNlY3JldCI6IjgyNjgwYmZlYzBmYTA4MzQ2YzFiMTBkMzBhM2UzZDRhIiwiZXhwIjoxNjk5MjEwMDYxLCJpc3MiOiJibG9nLXNlcnZpY2UifQ.RsZ3gFuKsO0vNMY8Fug8eTBMiT5b8ZISokR5-gIXfL8")
		println("this is res ", res)
		if err != nil {
			println("this is err ", err)

		}

		c.Next()
	}
}
