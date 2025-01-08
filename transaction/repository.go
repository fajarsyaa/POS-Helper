package transaction

import (
	"errors"
	"fmt"
	"slash/helper"
	"slash/product"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order Order) (Order, error)
	GetOrdersByUserId(UserId int) ([]Order, error)
	GetOrdersByUserIdAndOrderId(UserId int, OrderId string) (Order, error)
	PaymentNow(OrderId string) (Order, error)
	UpdateOrderByID(OrderId string, updatedItems OrderItem) (Order, error)
	DeleteOrderById(OrderId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
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

		err = trx.Save(&product).Error
		if err != nil {
			trx.Rollback()
			return order, fmt.Errorf("Not enough stock product id %d", order.OrderItems[i].ProductId)
		}

		order.OrderItems[i].Id = helper.GenerateRandomUUID2()
		if order.OrderItems[i].Id == "" {
			trx.Rollback()
			return order, fmt.Errorf("Failed to generate UUID for order item at index %d", i)
		}

		err = trx.Create(&order.OrderItems[i]).Error
		if err != nil {
			trx.Rollback()
			return order, err
		}
	}

	// Commit successful
	err = trx.Commit().Error
	if err != nil {
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

func (r *repository) GetOrdersByUserIdAndOrderId(UserId int, OrderId string) (Order, error) {
	var orders Order
	err := r.db.Preload("OrderItems", "id != ''").Preload("OrderItems.ItemDetail").Where("user_id = ? && id = ?", UserId, OrderId).Find(&orders).Error
	if err != nil {
		return Order{}, err
	}
	return orders, nil
}

func (r *repository) PaymentNow(OrderId string) (Order, error) {
	trx := r.db.Begin()

	// Defer rollback in case of panic
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	var order Order
	err := trx.Preload("OrderItems", "id != ''").Where("id = ?", OrderId).First(&order).Error
	if err != nil {
		trx.Rollback()
		return Order{}, err
	}

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

		fmt.Printf("%d - 1 %s", i, product)

		product.Stock -= order.OrderItems[i].Quantity
		product.UpdatedAt = time.Now()
		fmt.Printf("%d - 2 %s", i, product)
		err = trx.Save(&product).Error
		if err != nil {
			trx.Rollback()
			return order, err
		}
		fmt.Printf("%d - 3 %s", i, product)
	}
	order.Status = "done"
	err = trx.Save(&order).Error
	if err != nil {
		trx.Rollback()
		return Order{}, err
	}

	if err := trx.Commit().Error; err != nil {
		return Order{}, err
	}

	return order, nil
}

func (r *repository) UpdateOrderByID(OrderId string, updatedItems OrderItem) (Order, error) {
	trx := r.db.Begin()
	// Defer rollback in case of panic
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	var product product.Product
	err := trx.Where("id = ?", updatedItems.ProductId).First(&product).Error
	if err != nil {
		trx.Rollback()
		return Order{}, fmt.Errorf("Product Not Found")
	}

	if product.Stock < updatedItems.Quantity {
		trx.Rollback()
		return Order{}, fmt.Errorf("Not enough stock for product")
	}

	var order Order
	err = trx.Preload("OrderItems", "product_id = ? && id != ?", updatedItems.ProductId, "").Where("id = ?", OrderId).First(&order).Error
	if err != nil {
		trx.Rollback()
		return Order{}, fmt.Errorf("Order Not Found")
	}

	if order.Status == "done" {
		trx.Rollback()
		return Order{}, fmt.Errorf("Can't Update This Data. Because Already Paid")
	}

	if order.OrderItems == nil {
		trx.Rollback()
		return Order{}, fmt.Errorf("OrderItems Not Found")
	}

	for i, item := range order.OrderItems {
		if item.ProductId == updatedItems.ProductId {
			order.OrderItems[i].Quantity = updatedItems.Quantity
			break
		}
	}

	order.UpdatedAt = time.Now()
	err = trx.Save(&order.OrderItems).Error
	if err != nil {
		trx.Rollback()
		return Order{}, fmt.Errorf("Failed to update order")
	}

	err = trx.Save(&order).Error
	if err != nil {
		trx.Rollback()
		return Order{}, fmt.Errorf("Failed to update order")
	}

	err = trx.Commit().Error
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (r *repository) DeleteOrderById(OrderId string) error {
	var order Order
	err := r.db.Preload("OrderItems", "id != ''").Where("id = ?", OrderId).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Order With ID %s Not Found", OrderId)
		}
		return fmt.Errorf("Order Not Found : %v", err)
	}

	err = r.db.Delete(&order).Error
	if err != nil {
		return fmt.Errorf("Failed Delete Order : %v", err)
	}

	return nil
}
