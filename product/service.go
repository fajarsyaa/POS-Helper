package product

import "fmt"

type Service interface {
	GetAllProducts() ([]Product, error)
	FindProducts(keyword string) ([]Product, error)
	FindProductsById(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAllProducts() ([]Product, error) {
	products, err := s.repository.GetAllProducts()
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Get Product Data Failed")
	}
	return products, nil
}

func (s *service) FindProducts(keyword string) ([]Product, error) {
	product, err := s.repository.FindProducts(keyword)
	if err != nil {
		fmt.Println(err.Error())
		return product, fmt.Errorf("Find Product Data Failed")
	}
	return product, nil
}

func (s *service) FindProductsById(id int) (Product, error) {
	product, err := s.repository.FindProductsById(id)
	if err != nil {
		fmt.Println(err.Error())
		return product, fmt.Errorf("Find Product Data Failed")
	}
	return product, nil
}
