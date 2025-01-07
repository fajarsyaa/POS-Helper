package transaction

import "time"

type Order struct {
	Id              string    `json:"id"`
	UserId          int       `json:"user_id"`
	Total           float64   `json:"total"`
	CustomerName    string    `json:"customer_name"`
	CustomerPhone   string    `json:"customer_phone"`
	CustomerAddress string    `json:"customer_address"`
	Status          string    `json:"status"`
	ExpiredAt       time.Time `json:"expired_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	Id        string    `json:"id"`
	OrderId   string    `json:"order_id"`
	ProductId int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
