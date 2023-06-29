package middlewares

import (
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func (m Middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		originUrl, _ := url.Parse(c.Request.Header.Get("Origin"))
		originHost := originUrl.Hostname()

		// Set dev CORS policy
		if os.Getenv("ENV") == "dev" && (originHost == "localhost" || originHost == "127.0.0.1") {
			setCORSHeaders(c, c.Request.Header.Get("Origin"))
		}

		// Set prod CORS policy
		if os.Getenv("ENV") == "prod" && originHost == os.Getenv("PROD_URL") {
			setCORSHeaders(c, c.Request.Header.Get("Origin"))
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func setCORSHeaders(c *gin.Context, origin string) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, PATCH, GET, POST, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
}
