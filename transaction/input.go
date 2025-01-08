package transaction

type OrderInput struct {
	ProductInput    []ProductInput `json:"products" binding:"required"`
	OrdersTotal     float64        `json:"orders_total" binding:"required"`
	CustomerName    string         `json:"customer_name" binding:"required"`
	CustomerPhone   string         `json:"customer_phone"`
	CustomerAddress string         `json:"customer_address"`
	UserID          int
}

type ProductInput struct {
	ProductIId int `json:"products_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required"`
}

type OrderDetailInput struct {
	Id     string `json:"order_id"`
	UserId int
}

type PaymentInput struct {
	Id     string `json:"order_id"`
	UserId int
}

type UpdateOrderInput struct {
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	OrderID   string `json:"order_id"`
}
