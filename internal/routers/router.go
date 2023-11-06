package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-programming-tour-book/blog-service/docs"
	// "github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/qiuxiaori/go-blog/internal/middleware"
	v1 "github.com/qiuxiaori/go-blog/internal/routers/api/v1"
	"github.com/qiuxiaori/go-blog/pkg/limiter"
)


var article = v1.NewArticle()
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/api/v1/auth",
	FillInterval: time.Second * 10,
	Capacity:     1,
	Quantum:      1,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.AccessLog())
	r.Use(middleware.RateLimiter(methodLimiters))
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/articles", article.Create)
		apiV1.POST("/auth", v1.GetAuth)
	}

	return r
}
