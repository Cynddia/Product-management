package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ProcessBase64Image(base64Image string, fileName string) (string, error) {
	if strings.HasPrefix(base64Image, "data:image/") {
		parts := strings.Split(base64Image, ",")
		if len(parts) == 2 {
			base64Image = parts[1]
		} else {
			return "", errors.New("invalid base64 image format")
		}
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", err
	}

	dir := "images_compressed"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("Failed to create uploads directory: %v", err)
			return "", err
		}
	}

	filePath := dir + "/" + fileName
	err = ioutil.WriteFile(filePath, imageData, 0644)
	if err != nil {
		log.Printf("Failed to save image: %v", err)
		return "", err
	}

	return filePath, nil
}

func SaveImageLocally(base64Image string, fileName string) (string, error) {
	decodedImage, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %v", err)
	}

	filePath := fmt.Sprintf("uploads/%s", fileName)
	err = ioutil.WriteFile(filePath, decodedImage, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return filePath, nil
}
