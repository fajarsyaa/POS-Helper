package transaction

import "time"

type TransactionFormatter struct {
	OrderID     string    `json:"order_id"`
	ExpiredTime time.Time `json:"expired_time"`
}

type OrderFormatter struct {
	Order
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

func FormatterOrderResponse(order Order) OrderFormatter {
	formattedOrder := OrderFormatter{
		Order: order,
	}
	return formattedOrder
}
