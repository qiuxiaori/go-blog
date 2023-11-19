package qiurouter

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

type Payload struct {
	Param interface{}
	Body  interface{}
}

type Router struct {
	Method  string
	Url     string
	Handle  func(c *gin.Context)
	Payload Payload
}

type QiuRouter struct {
	Routers []Router
	R       *gin.Engine
}

func New(R *gin.Engine) *QiuRouter {
	return &QiuRouter{R: R}
}

func (r *QiuRouter) Register(router Router) {
	r.Routers = append(r.Routers, router)
	r.R.Handle(router.Method, router.Url, Validator(router.Payload.Body), router.Handle)
}

func SetPayload(Payload Payload) gin.HandlerFunc {
	return func(c *gin.Context) {
		if Payload.Param != nil {
			c.Set("Payload.Param", &Payload.Param)
			println("this is  Payload.Param ", &Payload.Param)
		}
		// c.Set("payload", Payload)
		// c.Set("payload", Payload)
		// c.Set("payload", Payload)
	}
}

func Validator(inter interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// param, exists := c.Get("Payload.Param")
		// if exists {
		// println("this is exists", exists, &param, param)
		valid, _ := app.BindAndValid(c, reflect.ValueOf(inter))
		println("this is inter1", inter, valid)
		if !valid {
			println("this is inter", inter, valid)
			// c.Abort()
			// return
		}
		// }

		c.Next()
	}
}
