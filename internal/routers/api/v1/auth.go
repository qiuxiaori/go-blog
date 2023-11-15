package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/internal/service"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

func GetAuth(c *gin.Context) {
	token, _ := app.GenerateToken("1212", "11212")
	app.NewResponse(c).ToResponse(token)
}

func GetUser(c *gin.Context) {
	auth, _ := service.Auth{}.Get("1212", "11212")
	app.NewResponse(c).ToResponse(auth)
}
