package models

import (
	"errors"
	"sync"
)

type Product struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImagePaths  []string `json:"image_paths"`
	Price       float64  `json:"price"`
}

type ProductInput struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageURLs   []string `json:"image_urls"`
	Price       float64  `json:"price"`
}

var (
	products         []Product
	productIDCounter int
	mu               sync.Mutex
)

func CreateProduct(p *Product) error {
	mu.Lock()
	defer mu.Unlock()

	productIDCounter++
	p.ID = productIDCounter

	products = append(products, *p)
	return nil
}

func GetProductByID(id int) (Product, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return Product{}, errors.New("product not found")
}

func GetProducts() ([]Product, error) {
	mu.Lock()
	defer mu.Unlock()

	return products, nil
}

func UpdateProduct(id int, updatedProduct ProductInput) (Product, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, p := range products {
		if p.ID == id {
			products[i].Title = updatedProduct.Title
			products[i].Description = updatedProduct.Description
			products[i].ImagePaths = updatedProduct.ImageURLs
			return products[i], nil
		}
	}

	return Product{}, errors.New("product not found")
}

func DeleteProduct(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}
