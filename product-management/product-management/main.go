package main

import (
	"log"
	"product-management/controllers"
	"product-management/services"
	"product-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	// Initialize Redis,
	utils.InitRedis()
	utils.InitLogger()

	db, err := utils.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	router.Static("/uploads", "./uploads")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	imageProcessor := &services.ImageService{}

	go func() {
		err := utils.ConsumeMessages(conn, imageProcessor, db)
		if err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}()

	router.Use(utils.LogRequests)

	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProductByID)
	router.GET("/products", controllers.GetProducts)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the Gin server: %v", err)
	}
}
