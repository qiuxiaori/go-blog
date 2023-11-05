package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/internal/middleware"
	v1 "github.com/qiuxiaori/go-blog/src/router/api/v1"
)

var article = v1.NewArticle()

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/articles", article.Create)
		apiV1.POST("/getToken", v1.GetAuth)
	}

	return r
}
