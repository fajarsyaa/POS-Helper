package product

import (
	"fmt"
	"strings"
	"time"
)

type ProductFormatter struct {
	Id        int    `json:"id"`
	SKU       string `json:"sku"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
	Price     string `json:"price"`
	CreatedAt string `json:"created_at"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	Image     string `json:"image"`
}

func FormatterProductResponses(products []Product) []ProductFormatter {
	var formattedProducts []ProductFormatter

	for _, product := range products {
		formatter := ProductFormatter{
			Id:        product.Id,
			SKU:       product.SKU,
			Name:      product.Name,
			Size:      product.Size,
			Color:     product.Color,
			Image:     product.Image,
			Stock:     product.Stock,
			Price:     formatPrice(product.Price),
			CreatedAt: formatDate(product.CreatedAt),
		}
		formattedProducts = append(formattedProducts, formatter)
	}

	return formattedProducts
}

func FormatterProductResponse(product Product) ProductFormatter {
	formatter := ProductFormatter{
		Id:        product.Id,
		SKU:       product.SKU,
		Name:      product.Name,
		Size:      product.Size,
		Color:     product.Color,
		Image:     product.Image,
		Stock:     product.Stock,
		Price:     formatPrice(product.Price),
		CreatedAt: formatDate(product.CreatedAt),
	}
	return formatter
}

func formatPrice(price float64) string {
	return fmt.Sprintf("%s", formatWithCommas(int(price)))
}

func formatWithCommas(num int) string {
	str := fmt.Sprintf("%d", num)
	n := len(str)
	if n <= 3 {
		return str
	}

	var result strings.Builder
	for i, digit := range str {
		if (n-i)%3 == 0 && i != 0 {
			result.WriteRune('.')
		}
		result.WriteRune(digit)
	}
	return result.String()
}

func formatDate(t time.Time) string {
	return t.Format("02-01-2006")
}
