package product

import "time"

type Product struct {
	Id        int
	SKU       string
	Name      string
	Size      string
	Color     string
	Image     string
	Price     float64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
