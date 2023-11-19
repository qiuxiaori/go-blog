package qrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/internal/qrouters/user"
	"github.com/qiuxiaori/go-blog/pkg/qiurouter"
)

func Init(r *gin.Engine) {
	routers := qiurouter.New(r)

	// apiV1 := r.Group("/api/v1")
	routers.Register(user.Create("POST", "/user"))
}
