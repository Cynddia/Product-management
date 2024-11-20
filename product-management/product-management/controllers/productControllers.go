package controllers

import (
	"log"
	"net/http"
	"strconv"

	"product-management/models"
	"product-management/services"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productInput models.ProductInput

	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var savedImagePaths []string
	for index, base64Image := range productInput.ImageURLs {
		fileName := "product_image_" + strconv.Itoa(index+1) + ".jpg"
		filePath, err := services.ProcessBase64Image(base64Image, fileName)
		if err != nil {
			log.Printf("Failed to process image %d: %v", index+1, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process image"})
			return
		}
		savedImagePaths = append(savedImagePaths, filePath)
	}

	product := models.Product{
		Title:       productInput.Title,
		Description: productInput.Description,
		ImagePaths:  savedImagePaths,
	}

	if err := models.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := models.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProductPrice(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := models.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": product.ID, "price": product.Price})
}

func GetProducts(c *gin.Context) {
	products, err := models.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
