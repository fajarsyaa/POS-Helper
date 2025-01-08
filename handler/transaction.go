package handler

import (
	"net/http"
	"slash/helper"
	"slash/transaction"
	"time"

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

	var expiredAt time.Time
	if order.ExpiredAt != nil {
		expiredAt = *order.ExpiredAt
	} else {
		expiredAt = time.Time{}
	}

	formatResponse := transaction.FormatterTRXResponse(order.Id, expiredAt)
	response := helper.ResponseMessage("Order Created", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) GetOrdersByUserId(ctx *gin.Context) {
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
	}

	orders, err := h.service.GetOrdersByUserId(int(userId))
	if err != nil {
		listErr := gin.H{"errors": err.Error()}
		response := helper.ResponseMessage("Create Failed", "Failed", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := transaction.FormatterOrderResponses(orders)
	response := helper.ResponseMessage("Order Created", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) GetOrdersByUserIdAndOrderId(ctx *gin.Context) {
	var input transaction.OrderDetailInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Get Data Failed", "bad request", http.StatusBadRequest, listErr)
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
		input.UserId = int(userId)
	}

	order, err := h.service.GetOrdersByUserIdAndOrderId(input.UserId, input.Id)
	if err != nil {
		listErr := gin.H{"errors": err.Error()}
		response := helper.ResponseMessage("Get Data Failed", "Failed", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := transaction.FormatterOrderResponse(order)
	response := helper.ResponseMessage("Get Data Failed", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) PaymentNow(ctx *gin.Context) {
	var input transaction.PaymentInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Payment Failed", "bad request", http.StatusBadRequest, listErr)
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
		input.UserId = int(userId)
	}

	order, err := h.service.PaymentNow(input.UserId, input.Id)
	if err != nil {
		listErr := gin.H{"errors": err.Error()}
		response := helper.ResponseMessage("Payment Failed", "Failed", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := transaction.FormatterPaymentResponse(order)
	response := helper.ResponseMessage("Payment Success", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) UpdateOrderById(ctx *gin.Context) {
	var input transaction.UpdateOrderInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Update Failed", "bad request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := h.service.UpdateOrderByID(input)
	if err != nil {
		listErr := gin.H{"errors": err.Error()}
		response := helper.ResponseMessage("Update Failed", "Failed", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := transaction.FormatterUpdateOrderResponse(order)
	response := helper.ResponseMessage("Update Success", "success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}
