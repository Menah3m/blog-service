package middleware

import (
	"github.com/gin-gonic/gin"
	"menah3m/blog-service/pkg/app"
	"menah3m/blog-service/pkg/errcode"
	"menah3m/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()

	}
}
