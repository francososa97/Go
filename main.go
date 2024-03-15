package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/francososa97/product-api/controller"
	"github.com/francososa97/product-api/repository"
	"github.com/francososa97/product-api/service"
)

func main() {
	router := mux.NewRouter()

	productRepository := repository.NewMongoRepository("mongodb://localhost:27017", "products")
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router.HandleFunc("/products", productController.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productController.GetProductByID).Methods("GET")
	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productController.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
