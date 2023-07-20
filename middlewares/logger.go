package middlewares

import (
	"backend/models"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (m Middleware) Logger() gin.HandlerFunc {
	file, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		// Before processing the request
		start := time.Now()

		// Process the request
		c.Next()

		// After processing the request
		end := time.Now()
		latency := end.Sub(start)
		method := c.Request.Method
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		// if ip is already stored in db, use it, otherwise get it from request
		ip := &models.IP{IP: clientIP}
		ip, err := ip.FirstOrCreate(m.db)
		if err != nil {
			logger.Println("Error occurred during IP geolocation: ", err)
			return
		}

		log := models.Log{
			Path:    path,
			Method:  method,
			Status:  statusCode,
			Latency: latency.Milliseconds(),
			IPID:    ip.ID,
		}

		m.db.Create(&log)

		// Log base request info
		logger.Printf("ClientIP: %s Method: %s Path: %s Status: %d Latency: %v", clientIP, method, path, statusCode, latency)

		// Log error if one occurred
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.Printf("Error in %v: %s", e.Meta, e.Err)
			}
		}
	}
}
