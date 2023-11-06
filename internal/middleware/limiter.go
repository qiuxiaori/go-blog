package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		bucket, ok := l.GetBucket(key)
		println("========= this is key", key, bucket, ok)
		if ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				// response := app.NewResponse(errcode.TooMandRequests)
				println("too may requests")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}