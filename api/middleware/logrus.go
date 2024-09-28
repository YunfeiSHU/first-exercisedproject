package middleware

import (
	"fmt"
	"gin-jwt-gorm/internal/logutil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logutil.NewLogger()

		timeFormat := "02/Jan/2006:15:04:05 -0700"
		hostname, _ := os.Hostname()
		path := c.Request.URL.Path
		start := time.Now()
		//中间件中的c.Next() 等待下一个中间件/处理方法 完成后，才执行之后的代码
		// logger 中间件的相关信息，需要在 执行完 指定路由的处理方法后，才运行
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		logger.WithFields(logrus.Fields{
			"hostname":          hostname,
			"status_code":       statusCode,
			"latency":           latency,
			"client_ip":         clientIP,
			"client_user_agent": clientUserAgent,
			"method":            c.Request.Method,
			"path":              path,
			"referer":           referer,
			"data_length":       dataLength,
		})

		if len(c.Errors) > 0 {
			logger.Error(c.Errors.String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if c.Writer.Status() > 499 {
				logger.Error(msg)
			} else if c.Writer.Status() > 399 {
				logger.Warn(msg)
			} else {
				logger.Info(msg)
			}
		}
	}
}
