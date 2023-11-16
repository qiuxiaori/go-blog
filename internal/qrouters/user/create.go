package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/internal/service"
	"github.com/qiuxiaori/go-blog/pkg/app"
	"github.com/qiuxiaori/go-blog/pkg/qiurouter"
)

func Create() qiurouter.Router {
	return qiurouter.Router{
		Method: "POST",
		Url:    "/create",
		Handle: func(c *gin.Context) {
			auth, _ := service.GetAuth("1212", "11212")
			app.NewResponse(c).ToResponse(auth)
		},
	}
}

// func GetUser(c *gin.Context) {
// 	auth, _ := service.GetAuth("1212", "11212")
// 	app.NewResponse(c).ToResponse(auth)
// }
