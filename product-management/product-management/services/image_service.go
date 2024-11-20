package services

import (
	"database/sql"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	_ "image/png"
)

type ImageService struct{}

func (s *ImageService) ProcessImage(imagePath string) (string, error) {
	imgFile, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Failed to open image: %v", err)
		return "", err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Printf("Failed to decode image: %v", err)
		return "", err
	}

	compressedPath := fmt.Sprintf("./compressed_images/%s_compressed.jpg", imagePath)

	if _, err := os.Stat("./compressed_images"); os.IsNotExist(err) {
		if err := os.Mkdir("./compressed_images", 0755); err != nil {
			log.Printf("Failed to create directory: %v", err)
			return "", err
		}
	}

	outFile, err := os.Create(compressedPath)
	if err != nil {
		log.Printf("Failed to create compressed image file: %v", err)
		return "", err
	}
	defer outFile.Close()

	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 75}) // Adjust quality as needed

	return compressedPath, nil
}

func (s *ImageService) UpdateProductWithImage(productID string, imagePath string, db *sql.DB) error {
	query := `UPDATE products SET compressed_product_images = $1 WHERE id = $2`
	if _, err := db.Exec(query, imagePath, productID); err != nil {
		log.Printf("Failed to update product in database: %v", err)
		return err
	}
	return nil
}
