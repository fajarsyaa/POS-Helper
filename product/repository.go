package product

import "gorm.io/gorm"

type Repository interface {
	GetAllProducts() ([]Product, error)
	FindProducts(keyword string) ([]Product, error)
	FindProductsById(id int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllProducts() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindProducts(keyword string) ([]Product, error) {
	var products []Product
	err := r.db.Where("name LIKE ?", "%"+keyword+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindProductsById(id int) ([]Product, error) {
	var products []Product
	err := r.db.Where("id = ?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
