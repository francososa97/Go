# Products API in Go

This is a simple API for managing products developed in Go. It allows performing CRUD (Create, Read, Update, Delete) operations on products.

## Features

- Create a new product with name, description, price, and stock quantity.
- Get the list of all products.
- Get a specific product by its ID.
- Update the information of an existing product.
- Delete a product by its ID.

## Usage

### Installation

1. Clone this repository:

```bash
git clone https://github.com/tu_usuario/api-productos-go.git
```

2. Change directory to the project:

```bash
cd api-products-go
```

3. Install dependencies:

```bash
go mod tidy
```
## Configuration
Configure the database in the config/config.go file. Update the connection values according to your environment:
```
const (
    DBHost     = "localhost"
    DBPort     = "5432"
    DBUser     = "tu_usuario"
    DBPassword = "tu_contraseña"
    DBName     = "nombre_base_de_datos"
)
```

## Execution
Compile and run the application:
```
./api-productos-go
La API estará disponible en http://localhost:8080.
```


## Endpoints
- GET /products: Get the list of all products.
- GET /products/{id}: Get a specific product by its ID.
- POST /products: Create a new product.
- PUT /products/{id}: Update the information of an existing product.
- DELETE /products/{id}: Delete a product by its ID.
