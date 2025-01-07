package product

import (
	"fmt"
	"strings"
	"time"
)

type ProductFormatter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       string `json:"price"`
	CreatedAt   string `json:"created_at"`
}

func FormatterProductResponse(products []Product) []ProductFormatter {
	var formattedProducts []ProductFormatter

	for _, product := range products {
		formatter := ProductFormatter{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       formatPrice(product.Price),
			CreatedAt:   formatDate(product.CreatedAt),
		}
		formattedProducts = append(formattedProducts, formatter)
	}

	return formattedProducts
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
