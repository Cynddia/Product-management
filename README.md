# Overview
This Product Management System is designed to handle the creation, storage, and processing of product-related data, including high-resolution images. The system is built using a modular approach with separate services for image processing and product management, leveraging technologies such as RabbitMQ, Redis, GORM, and Gin. This README will explain the architectural choices, setup instructions, and assumptions made during the development process.

# Architectural Choices
## 1. Microservices Architecture

RabbitMQ: RabbitMQ is used as a message broker to handle asynchronous tasks such as image compression. Instead of blocking the main application thread while images are processed, product creation requests enqueue image processing jobs. These jobs are consumed by a separate service, improving performance and responsiveness.
Redis: Redis is utilized for caching frequently accessed product data. This helps in reducing the number of database queries, which improves the response time for read-heavy operations such as retrieving product details.
GORM: GORM is used as the ORM (Object Relational Mapping) tool for interacting with the PostgreSQL database. It simplifies the database operations such as querying, creating, updating, and deleting records while providing a higher-level abstraction over SQL queries.
Gin Framework: The Gin web framework is chosen for building the API server due to its lightweight, high-performance nature and easy-to-use routing capabilities. It supports RESTful APIs and is ideal for handling HTTP requests and responses.
## 2. Asynchronous Processing
The system offloads the image compression task into an asynchronous process using RabbitMQ. When a new product is created, the system sends the product's image URLs to a RabbitMQ queue. The image processing service, which consumes messages from this queue, handles the compression and updates the database with the compressed image paths. This ensures that the user-facing API remains fast, and long-running tasks are handled in the background.
## 3. Database Design
A PostgreSQL database is used to store product information. Each product has a set of properties such as name, description, price, and associated images. The compressed images are stored in the database, and their paths are updated once processed.

# Setup Instructions
Prerequisites
The required softwares are stated as follows:

Go (Golang): The programming language used for the backend service.
PostgreSQL: The database where product data will be stored.
Redis: The caching layer for improving read performance.
RabbitMQ: The message broker for asynchronous image processing.
Go Modules: Dependency management in Go.

# Step-by-Step Setup

Install Dependencies Run the following command to install the required Go dependencies:

Set Up PostgreSQL Database to have running postgress instance.

Create a new database called product_management and update the connection settings.

Set Up Redis and rabbitmq

# Running the Application

Once the services are set up, run the main application:

go run main.go


Test the API endpoints using Postman . The following endpoints are available:
POST /products: Creates a new product and enqueues image URLs for processing.
GET /products/{id}: Retrieves a specific product by ID.
GET /products: Retrieves a list of all products.


# Conclusion

This product management system demonstrates how microservices, asynchronous processing, and caching can be integrated into a cohesive solution to manage products and their associated images efficiently. With RabbitMQ handling the image compression in the background and Redis enhancing read performance, this system is both scalable and responsive, making it suitable for high-volume e-commerce applications.
