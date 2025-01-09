package transaction

import (
	"fmt"
	"slash/helper"
	"time"
)

type Service interface {
	CreateOrder(orderInput OrderInput) (Order, error)
	GetOrdersByUserId(UserId int) ([]Order, error)
	GetOrdersByUserIdAndOrderId(UserIdintint int, OrderId string) (Order, error)
	PaymentNow(UserId int, OrderId string) (Order, error)
	UpdateOrderByID(Input UpdateOrderInput) (Order, error)
	DeleteOrderById(OrderId string) error
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
		ExpiredAt:       time.Now().Add(1 * time.Hour),
		Status:          "pending",
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
		fmt.Println(err.Error())
		return Order{}, fmt.Errorf("Create Order Failed")
	}

	return savedOrder, nil
}

func (s *service) GetOrdersByUserId(UserId int) ([]Order, error) {
	orders, err := s.repository.GetOrdersByUserId(UserId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Data Not Found")
	}
	return orders, nil
}

func (s *service) GetOrdersByUserIdAndOrderId(UserId int, OrderId string) (Order, error) {
	order, err := s.repository.GetOrdersByUserIdAndOrderId(UserId, OrderId)
	if err != nil {
		fmt.Println(err.Error())
		return Order{}, fmt.Errorf("Data Not Found")
	}

	return order, nil
}

func (s *service) PaymentNow(UserId int, OrderId string) (Order, error) {
	exist, err := s.repository.GetOrdersByUserIdAndOrderId(UserId, OrderId)
	if err != nil {
		fmt.Println(err.Error())
		return Order{}, fmt.Errorf("Data Not Found")
	}

	if exist.Id == "" {
		return Order{}, fmt.Errorf("Data Not Found")
	}

	if exist.Status == "done" {
		return Order{}, fmt.Errorf("Bill Already Paid")
	}

	if exist.ExpiredAt.Before(time.Now()) {
		return Order{}, fmt.Errorf("Expired Billing Payment")
	}

	order, err := s.repository.PaymentNow(OrderId)
	if err != nil {
		fmt.Println(err.Error())
		return Order{}, fmt.Errorf("Payment Failed")
	}
	return order, nil
}

func (s *service) UpdateOrderByID(Input UpdateOrderInput) (Order, error) {
	if Input.OrderID == "" {
		return Order{}, fmt.Errorf("Orders Not Found")
	}

	orderItems := OrderItem{
		Quantity:  Input.Quantity,
		ProductId: Input.ProductID,
	}

	order, err := s.repository.UpdateOrderByID(Input.OrderID, orderItems)
	if err != nil {
		fmt.Println(err.Error())
		return Order{}, fmt.Errorf("Update Failed")
	}

	return order, nil
}

func (s *service) DeleteOrderById(OrderId string) error {
	err := s.repository.DeleteOrderById(OrderId)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Delete Failed")
	}

	return nil
}
