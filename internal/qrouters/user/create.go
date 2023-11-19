package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/internal/service"
	"github.com/qiuxiaori/go-blog/pkg/app"
	"github.com/qiuxiaori/go-blog/pkg/qiurouter"
)

func Create(method string, url string) qiurouter.Router {
	return qiurouter.Router{
		Method: method,
		Url:    url,
		Payload: qiurouter.Payload{
			Body: service.CreateUserReq{},
		},
		Handle: func(c *gin.Context) {
			var param service.CreateUserReq
			println("this is Handle")
			c.ShouldBindJSON(&param)
			service.CreateUser(&param)
			app.NewResponse(c).ToResponse(param)
		},
	}
}
