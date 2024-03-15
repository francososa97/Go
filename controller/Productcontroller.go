package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/francososa97/product-api/repository"
	"github.com/francososa97/product-api/service"
)

type ProductController struct {
	Service service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	sortByPriceAsc := r.URL.Query().Get("sortByPriceAsc") == "true"
	products, err := c.Service.GetAllProducts(sortByPriceAsc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	product, err := c.Service.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product repository.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.Service.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var product repository.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateProduct(id, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := c.Service.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
