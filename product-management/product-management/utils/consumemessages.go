package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"product-management/models"
	"product-management/services"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func ConsumeMessages(conn *amqp.Connection, imageProcessor *services.ImageService, db *gorm.DB) error {
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	queueName := "imageProcessingQueue"
	_, err = ch.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare the queue: %v", err)
	}

	msgs, err := ch.Consume(
		queueName,
		"",    // Consumer name
		true,  // Auto-ack
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	for d := range msgs {
		var message map[string]interface{}
		if err := json.Unmarshal(d.Body, &message); err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		productID := message["product_id"].(float64)
		imageURLs := message["image_urls"].([]interface{})

		for _, url := range imageURLs {
			imagePath := url.(string)

			log.Printf("Processing image: %s", imagePath)

			compressedImagePath, err := imageProcessor.ProcessImage(imagePath)
			if err != nil {
				log.Printf("Error processing image: %v", err)
				continue
			}

			if err := db.Model(&models.Product{}).Where("id = ?", productID).Update("compressed_product_images", compressedImagePath).Error; err != nil {
				log.Printf("Error updating product with image: %v", err)
			}
		}
	}
	return nil
}
