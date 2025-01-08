package transaction

import (
	"time"
)

type TransactionFormatter struct {
	OrderID     string    `json:"order_id"`
	ExpiredTime time.Time `json:"expired_time"`
}

type OrderFormatter struct {
	Order
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

func FormatterOrderResponses(orders []Order) []OrderFormatter {
	var formattedOrders []OrderFormatter

	for _, order := range orders {
		formattedOrder := OrderFormatter{
			Order: order,
		}
		formattedOrders = append(formattedOrders, formattedOrder)
	}

	return formattedOrders
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

func FormatterOrderResponse(order Order) OrderFormatter {
	formattedOrder := OrderFormatter{
		Order: order,
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
