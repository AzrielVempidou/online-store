# Online Store Backend

This is a RESTful API for an online store built with Go. The backend supports user registration, authentication, product listing, cart management, and checkout.

## Features

- User Registration
- User Login with JWT Authentication
- View Products by Category
- Add Products to Shopping Cart
- View Shopping Cart
- Delete Products from Shopping Cart
- Checkout

## Installation

1. **Clone the repository:**
  ```bash
  git clone https://github.com/AzrielVempidou/online-store.git
  cd online-store
  ```
2. **Install dependencies:**
  ```bash
    go mod tidy
  ```

3. **Run the application:**
  ```bash
    go run main.go
    The server will start on port 8080.
  ```

# API Endpoints
## Authentication
- Register
- URL: /register
- Method: POST
- Request Body:
```json
{
  "username": "testuser",
  "password": "password123",
  "email": "testuser@example.com"
}
```
- Response:
```json
{
  "message": "User registered successfully"  
}
  ```

## Login
- URL: /login
- Method: POST
- Request Body:
```json
{
  "username": "testuser",
  "password": "password123"
}
```

- Response:
```json
{
  "token": "your.jwt.token"
}
```

## Products
- Get Products
- URL: /products
- Method: GET
- Query Parameters: category (optional)
- Response:
```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-06-25T12:00:00Z",
    "UpdatedAt": "2024-06-25T12:00:00Z",
    "DeletedAt": null,
    "name": "Product 1",
    "category": "Category 1",
    "price": 100.0,
    "description": "Description for Product 1"
  },
  {
    "ID": 2,
    "CreatedAt": "2024-06-25T12:00:00Z",
    "UpdatedAt": "2024-06-25T12:00:00Z",
    "DeletedAt": null,
    "name": "Product 2",
    "category": "Category 2",
    "price": 200.0,
    "description": "Description for Product 2"
  }
]
```

## Get Products by Category
- Get Products by Category
- URL: /products/{category}
- Method: GET
- Query Parameters: category (optional)
- Response:
```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-06-25T12:00:00Z",
    "UpdatedAt": "2024-06-25T12:00:00Z",
    "DeletedAt": null,
    "name": "Product 1",
    "category": "Category 1",
    "price": 100.0,
    "description": "Description for Product 1"
  }
]
```

## Cart
- Get Cart
- URL: /cart
- Method: GET
- Headers:
  - Authorization: Bearer your.jwt.token

```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-06-25T12:00:00Z",
    "UpdatedAt": "2024-06-25T12:00:00Z",
    "DeletedAt": null,
    "user_id": 1,
    "product_id": 1,
    "quantity": 2,
    "product": {
      "ID": 1,
      "CreatedAt": "2024-06-25T12:00:00Z",
      "UpdatedAt": "2024-06-25T12:00:00Z",
      "DeletedAt": null,
      "name": "Product 1",
      "category": "Category 1",
      "price": 100.0,
      "description": "Description for Product 1"
    }
  }
]
```

## Add to Cart
- URL: /cart
- Method: POST
- Headers:
  - Authorization: Bearer your.jwt.token

- Request Body:
```json
{
  "product_id": 1,
  "quantity": 2
}
```
- Response :
```json
{
  "ID": 1,
  "CreatedAt": "2024-06-25T12:00:00Z",
  "UpdatedAt": "2024-06-25T12:00:00Z",
  "DeletedAt": null,
  "user_id": 1,
  "product_id": 1,
  "quantity": 2
}
```

## Delete from Cart
- URL: /cart/{productId}
- Method: DELETE
- Headers:
  - Authorization: Bearer your.jwt.token
- Request Body :
```json
{
  "product_id": 1
}
```
- Response:
```json
{
  "message": "Item deleted from cart"
}
```

## Checkout

- URL: /checkout
- Method: POST
- Headers:
  - Authorization: Bearer your.jwt.token
- Response:
```json
{
  "message": "Checkout successful"
}
```
