package user

import (
	"fmt"

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
			Body: service.UserCreateReq{},
		},
		Handle: func(c *gin.Context) {
			var param service.UserCreateReq
			if err := c.ShouldBindJSON(&param); err != nil {
				println("this is err %s", &err)
				fmt.Println(err)
				app.NewResponse(c).ToResponse("err")
				return
			}
			println("this is param", param.Name)
			if err := service.CreateUser(&param); err != nil {
				fmt.Println(err)
				app.NewResponse(c).ToResponse("err")
			}
			app.NewResponse(c).ToResponse("")
		},
	}
}
