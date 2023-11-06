package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

func GetAuth(c *gin.Context) {
	token, _ := app.GenerateToken("1212", "11212")
	app.NewResponse(c).ToResponse(token)
}
