package transaction

import (
	"time"
)

type TransactionFormatter struct {
	OrderID     string    `json:"order_id"`
	ExpiredTime time.Time `json:"expired_time"`
}

type OrderDetailFormatter struct {
	ID              string               `json:"id"`
	UserID          int                  `json:"user_id"`
	Total           float64              `json:"total"`
	CustomerName    string               `json:"customer_name"`
	CustomerPhone   string               `json:"customer_phone"`
	CustomerAddress string               `json:"customer_address"`
	Status          string               `json:"status"`
	ExpiredAt       time.Time            `json:"expired_at"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	OrderItems      []OrderItemFormatter `json:"order_items"`
}

type OrderItemFormatter struct {
	ID        string    `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	SKU       string    `json:"SKU"`
	Name      string    `json:"Name"`
	Size      string    `json:"Size"`
	Color     string    `json:"Color"`
	Image     string    `json:"Image"`
	Price     float64   `json:"Price"`
	Stock     int       `json:"Stock"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllOrder struct {
	OrderID     string
	Name        string
	ExpiredTime time.Time
	Status      string
}

type PaymentFormatter struct {
	OrderId string  `json:"order_id"`
	Status  string  `json:"status"`
	Total   float64 `json:"total"`
}

type UpdateOrderFormatter struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status_order"`
}

func FormatterTRXResponse(id string, exp time.Time) TransactionFormatter {
	return TransactionFormatter{OrderID: id, ExpiredTime: exp}
}

func FormatterAllOrderResponses(orders []Order) []GetAllOrder {
	var formattedOrders []GetAllOrder

	for _, order := range orders {
		formattedOrder := GetAllOrder{
			OrderID:     order.Id,
			Name:        order.CustomerName,
			ExpiredTime: order.ExpiredAt,
			Status:      order.Status,
		}
		formattedOrders = append(formattedOrders, formattedOrder)
	}

	return formattedOrders
}

func FormatterOrderResponse(order Order) OrderDetailFormatter {
	var orderItems []OrderItemFormatter
	for _, item := range order.OrderItems {
		newItem := OrderItemFormatter{
			ID:        item.Id,
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
			SKU:       item.ItemDetail.SKU,
			Name:      item.ItemDetail.Name,
			Size:      item.ItemDetail.Size,
			Color:     item.ItemDetail.Color,
			Image:     item.ItemDetail.Image,
			Price:     item.ItemDetail.Price,
			Stock:     item.ItemDetail.Stock,
			UpdatedAt: item.UpdatedAt,
		}

		orderItems = append(orderItems, newItem)
	}

	formattedOrder := OrderDetailFormatter{
		ID:              order.Id,
		UserID:          order.UserId,
		Total:           order.Total,
		CustomerName:    order.CustomerName,
		CustomerPhone:   order.CustomerPhone,
		CustomerAddress: order.CustomerAddress,
		Status:          order.Status,
		ExpiredAt:       order.ExpiredAt,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
		OrderItems:      orderItems,
	}
	return formattedOrder
}

func FormatterPaymentResponse(order Order) PaymentFormatter {
	return PaymentFormatter{
		Status:  order.Status,
		Total:   order.Total,
		OrderId: order.Id,
	}
}

func FormatterUpdateOrderResponse(order Order) UpdateOrderFormatter {
	return UpdateOrderFormatter{
		Status:  order.Status,
		OrderId: order.Id,
	}
}
