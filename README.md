# Store API

## Description
Store API is a RESTful API built with Go (Golang) for managing store data, including products, users, and SKUs (Stock Keeping Units). It provides endpoints for user registration, product management, and SKU management, allowing users to create, read, update, and delete resources.

## Features
- User authentication and management
- Product creation, retrieval, updating, and deletion
- SKU management for products
- Image upload functionality for user avatars and product images
- Swagger documentation for API endpoints

## Installation Guide

### Prerequisites
- Go (version 1.18 or higher)
- PostgreSQL database (version 13 or higher)
- Minio Object Storage

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/MociW/store-api-golang.git
   cd store-api-golang
   ```

2. Install dependencies:
    ```bash
    go mod tidy
    go get -u github.com/swaggo/swag/cmd/swag
    ```

3. Set up environment variables: `config.yaml`
    ```yml
    server:
     Host: "<YOUR_SERVER_HOST>"
     Port: "<YOUR_SERVER_PORT>"
     JWTSecretKey: "<YOUR_JWT_SECRET_KEY>"

   database:
     Host: "<DATABASE_HOST>"
     Port: "<DATABASE_PORT>"
     User: "<DATABASE_USERNAME>"
     Password: "<DATABASE_PASSWORD>"
     NameDB: "<DATABASE_NAME>"

   aws:
     Endpoint: "<AWS_ENDPOINT>"
     MinioAccessKey: "<MINIO_ACCESS_KEY>"
     MinioSecretKey: "<MINIO_SECRET_KEY>"
     UseSSL: true/false
    ```

4. Run the application:
    ```bash
    make start
    ```

5. Access Swagger documentation:
    ```bash
    https://{YOUR_SERVER_HOST:YOUR_SERVER_PORT}/swagger/index.html
    ```

## API Endpoints
### User Endpoints
- POST /users: Create a new user account
- POST /users/login: Authenticate a user and return access and refresh tokens
- GET /users/me: Retrieve the current user's information
- PUT /users/me: Update the current user's information
- POST /users/me/avatar: Upload a new avatar for the current user
- POST /users/me/password: Update the current user's password
- GET /users/{id}: Retrieve a user's information by ID

### Product Endpoints
- POST /products: Create a new product
- GET /products/{id}: Find a product by ID
- PUT /products/{id}: Update an existing product
- DELETE /products/{id}: Delete a product by ID
- GET /products/list: List all products for a user
- GET /products/search: Search for products by name or description

### SKU Endpoints
- POST /products/{id}/sku: Create a new SKU for a product
- GET /products/{id}/sku/{skuId}: Find a SKU by ID
- DELETE /products/{id}/sku/{skuId}: Delete a SKU by ID
- GET /products/{id}/sku/list: List all SKUs for a product

