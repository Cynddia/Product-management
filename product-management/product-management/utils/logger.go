package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {

	Log.SetFormatter(&logrus.JSONFormatter{})

	Log.SetLevel(logrus.DebugLevel)
}

func LogRequests(c *gin.Context) {

	startTime := time.Now()

	c.Next()

	duration := time.Since(startTime)

	Log.WithFields(logrus.Fields{
		"method":   c.Request.Method,
		"uri":      c.Request.RequestURI,
		"clientIP": c.ClientIP(),
		"status":   c.Writer.Status(),
		"duration": duration,
	}).Info("Request completed")
}
