package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/global"
	"github.com/qiuxiaori/go-blog/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Writer(p []byte) (int, error) {
	if n,err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{ c.Writer, bytes.NewBufferString("")}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields {
			"request": c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, negin_time: %d, end_time: %d",
		c.Request.Method,
		bodyWriter.Status(),
		beginTime,
		endTime,
	)

	}
}