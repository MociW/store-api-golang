# Store API

## Description
Store API is a RESTful API built with Go (Golang) for managing store data, including products, users, and SKUs (Stock Keeping Units). It provides endpoints for user registration, product management, and SKU management, allowing users to create, read, update, and delete resources.

## Features
- User authentication and management
- Product creation, retrieval, updating, and deletion
- SKU management for products
- Image upload functionality for user avatars and product images
- Swagger documentation for API endpoints

## Installation

### Prerequisites
- Go (version 1.16 or higher)
- PostgreSQL database
- AWS account (for image storage)

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/MociW/store-api-golang.git
   cd store-api-golang

2. Install dependencies:
    ```bash
    go mod tidy

3. Set up environment variables: `config.yaml`
    ```config.yaml
   