package service

import (
	"errors"

	"github.com/francososa97/product-api/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{Repository: repo}
}

func (s *ProductService) GetAllProducts(sortByPriceAsc bool) ([]repository.Product, error) {
	return s.Repository.GetAll(sortByPriceAsc)
}

func (s *ProductService) GetProductByID(id string) (*repository.Product, error) {
	return s.Repository.GetByID(id)
}

func (s *ProductService) CreateProduct(product *repository.Product) error {
	return s.Repository.Create(product)
}

func (s *ProductService) UpdateProduct(id string, product *repository.Product) error {
	_, err := s.Repository.GetByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.Repository.Update(id, product)
}

func (s *ProductService) DeleteProduct(id string) error {
	_, err := s.Repository.GetByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.Repository.Delete(id)
}
