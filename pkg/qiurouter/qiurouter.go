package qiurouter

import "github.com/gin-gonic/gin"

type Router struct {
	Method      string
	Url         string
	Handle      func(c *gin.Context)
	Middlewares []func() gin.HandlerFunc
	Payload interface{}
}

type QiuRouter struct {
	Routers []Router
	Gin     *gin.Engine
}

func New(Gin *gin.Engine) *QiuRouter {
	return &QiuRouter{Gin: Gin}
}

// func
func (r *QiuRouter) Register(router Router) {
	r.Routers = append(r.Routers, router)
	// api := r.Gin.Group()
	r.Gin.Handle(router.Method, router.Url, router.Handle, ...router.Middlewares)
}
