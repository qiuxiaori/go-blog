package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/internal/service"
	"github.com/qiuxiaori/go-blog/pkg/app"
	"github.com/qiuxiaori/go-blog/pkg/qiurouter"
)

func Detail(method string, url string) qiurouter.Router {
	return qiurouter.Router{
		Method: method,
		Url:    url,
		Payload: qiurouter.Payload{
			Body: service.UserCreateReq{},
		},
		Handle: func(c *gin.Context) {
			var param service.UserDetailReq
			c.ShouldBindJSON(&param)
			user, err := service.GetUser(&param)
			if err != nil {
				fmt.Println(err)
				app.NewResponse(c).ToResponse("err")
				return
			}
			app.NewResponse(c).ToResponse(user)
		},
	}
}
