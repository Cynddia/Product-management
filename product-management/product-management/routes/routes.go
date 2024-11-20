package routes

import (
	"product-management/controllers"

	"product-management/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		utils.Log.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"route":    c.FullPath(),
			"status":   c.Writer.Status(),
			"duration": duration.Seconds(),
			"ip":       c.ClientIP(),
		}).Info("Request processed")
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProductByID)
	router.GET("/products", controllers.GetProducts)

	return router
}
