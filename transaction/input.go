package transaction

type OrderInput struct {
	ProductInput    []ProductInput `json:"products" binding:"required"`
	OrdersTotal     float64        `json:"orders_total" binding:"required"`
	CustomerName    string         `json:"customer_name" binding:"required"`
	CustomerPhone   string         `json:"customer_phone"`
	CustomerAddress string         `json:"customer_phone"`
	UserID          int
}

type ProductInput struct {
	ProductIId int `json:"products_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required"`
}
