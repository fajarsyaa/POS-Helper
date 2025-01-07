package handler

import (
	"net/http"
	"slash/helper"
	"slash/transaction"

	"github.com/gin-gonic/gin"
)

type trxHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *trxHandler {
	return &trxHandler{service}
}

func (h *trxHandler) CreateOrder(ctx *gin.Context) {
	var input transaction.OrderInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Created Order Failed", "bad request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, ok := ctx.Get("userID")
	if !ok {
		listErr := gin.H{"errors": "User  ID not found in context"}
		response := helper.ResponseMessage("System Error", "User  ID not found", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	userId, ok := userID.(float64)
	if !ok {
		listErr := gin.H{"errors": "Internal Server Error"}
		response := helper.ResponseMessage("System Error", "Internal Server Error", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	} else {
		input.UserID = int(userId)
	}

	order, err := h.service.CreateOrder(input)
	if err != nil {
		listErr := gin.H{"errors": err.Error()}
		response := helper.ResponseMessage("Create Failed", "Failed", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := transaction.FormatterTRXResponse(order.Id, order.ExpiredAt)
	response := helper.ResponseMessage("Order Created", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}
