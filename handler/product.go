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
			response := helper.ResponseMessage("Failed", "Product Not Found", http.StatusNotFound, err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
	}

	formatResponse := product.FormatterProductResponses(products)
	response := helper.ResponseMessage("Success", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProductByName(ctx *gin.Context) {
	var input product.FindProductByNameInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Failed", "bad request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.FindProducts(input.Keyword)
	if err != nil {
		if err != nil {
			response := helper.ResponseMessage("Failed", "Product Not Found", http.StatusNotFound, err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
	}

	formatResponse := product.FormatterProductResponses(products)
	response := helper.ResponseMessage("Success", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProductById(ctx *gin.Context) {
	var input product.FindProductByIdInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Failed", "bad request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.FindProductsById(input.Id)
	if err != nil {
		if err != nil {
			response := helper.ResponseMessage("Failed", "Product Not Found", http.StatusNotFound, err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
	}

	formatResponse := product.FormatterProductResponse(products)
	response := helper.ResponseMessage("Success", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}
