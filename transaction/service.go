package transaction

import (
	"slash/helper"
	"time"
)

type Service interface {
	CreateOrder(orderInput OrderInput) (Order, error)
	GetOrdersByUserId(UserId int) ([]Order, error)
	GetOrdersByUserIdAndOrderId(UserIdintint int, OrderId string) (Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateOrder(orderInput OrderInput) (Order, error) {

	order := Order{
		Id:              helper.GenerateRandomUUID(),
		UserId:          orderInput.UserID,
		Total:           orderInput.OrdersTotal,
		CustomerName:    orderInput.CustomerName,
		CustomerPhone:   orderInput.CustomerPhone,
		CustomerAddress: orderInput.CustomerAddress,
		Status:          "pending",
		ExpiredAt:       time.Now().Add(1 * time.Hour),
		CreatedAt:       time.Now(),
	}

	var newOrderItems []OrderItem
	for _, productInput := range orderInput.ProductInput {
		orderItem := OrderItem{
			OrderId:   order.Id,
			ProductId: productInput.ProductIId,
			Quantity:  productInput.Quantity,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		newOrderItems = append(newOrderItems, orderItem)
	}

	order.OrderItems = newOrderItems
	savedOrder, err := s.repository.CreateOrder(order)
	if err != nil {
		return Order{}, err
	}

	return savedOrder, nil
}

func (s *service) GetOrdersByUserId(UserId int) ([]Order, error) {
	orders, err := s.repository.GetOrdersByUserId(UserId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *service) GetOrdersByUserIdAndOrderId(UserId int, OrderId string) (Order, error) {
	orders, err := s.repository.GetOrdersByUserIdAndOrderId(UserId, OrderId)
	if err != nil {
		return Order{}, err
	}
	return orders, nil
}
