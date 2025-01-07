package transaction

import (
	"fmt"
	"slash/helper"
	"slash/product"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order Order) (Order, error)
	GetOrdersByUserId(UserId int) ([]Order, error)
}

type repository struct {
	db          *gorm.DB
	productRepo product.Repository
}

func NewRepository(db *gorm.DB, productRepo product.Repository) *repository {
	return &repository{db: db, productRepo: productRepo}
}

func (r *repository) CreateOrder(order Order) (Order, error) {
	trx := r.db.Begin()

	// Defer rollback in case of panic
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	// Create order
	err := trx.Create(&order).Error
	if err != nil {
		trx.Rollback()
		return order, err
	}

	// Create order items
	for i := range order.OrderItems {
		var product product.Product

		err := trx.First(&product, order.OrderItems[i].ProductId).Error
		if err != nil {
			trx.Rollback()
			return order, fmt.Errorf("Not Found product id %d", order.OrderItems[i].ProductId)
		}

		if product.Stock < order.OrderItems[i].Quantity {
			trx.Rollback()
			return order, fmt.Errorf("Not enough stock product id %d", order.OrderItems[i].ProductId)
		}

		product.Stock -= order.OrderItems[i].Quantity
		product.UpdatedAt = time.Now()

		err = trx.Save(&product).Error
		if err != nil {
			trx.Rollback()
			return order, fmt.Errorf("Not enough stock product id %d", order.OrderItems[i].ProductId)
		}

		order.OrderItems[i].Id = helper.GenerateRandomUUID2()
		err = trx.Create(&order.OrderItems[i]).Error
		if err != nil {
			trx.Rollback()
			return order, err
		}

	}

	// Commit successful
	if err := trx.Commit().Error; err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) GetOrdersByUserId(UserId int) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ?", UserId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
