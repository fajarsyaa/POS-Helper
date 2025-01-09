package handler

import (
	"net/http"
	"slash/helper"
	"slash/product"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service: service}
}

func (h *productHandler) GetAllProduct(ctx *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		if err != nil {
			listErr := gin.H{"error": err.Error()}
			response := helper.ResponseMessage("Failed", "Product Not Found", http.StatusNotFound, listErr)
			ctx.JSON(http.StatusNotFound, response)
			return
		}
	}

	formatResponse := product.FormatterProductResponses(products)
	response := helper.ResponseMessage("Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProductByName(ctx *gin.Context) {
	var input product.FindProductByNameInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"error": errors}

		response := helper.ResponseMessage("Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.FindProducts(input.Keyword)
	if err != nil {
		listErr := gin.H{"error": err.Error()}
		response := helper.ResponseMessage("Failed", "Product Not Found", http.StatusNotFound, listErr)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if len(products) == 0 {
		listErr := gin.H{"error": "Product Not Found"}
		response := helper.ResponseMessage("Success", "Product Not Found", http.StatusOK, listErr)
		ctx.JSON(http.StatusOK, response)
		return
	}

	formatResponse := product.FormatterProductResponses(products)
	response := helper.ResponseMessage("Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProductById(ctx *gin.Context) {
	var input product.FindProductByIdInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"error": errors}

		response := helper.ResponseMessage("Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.FindProductsById(input.Id)
	if err != nil {
		if err != nil {
			listErr := gin.H{"error": err.Error()}
			response := helper.ResponseMessage("Failed", "Product Not Found", http.StatusNotFound, listErr)
			ctx.JSON(http.StatusNotFound, response)
			return
		}
	}

	if products.Id == 0 {
		listErr := gin.H{"error": "Product Not Found"}
		response := helper.ResponseMessage("Success", "Product Not Found", http.StatusOK, listErr)
		ctx.JSON(http.StatusOK, response)
		return
	}

	formatResponse := product.FormatterProductResponse(products)
	response := helper.ResponseMessage("Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}
