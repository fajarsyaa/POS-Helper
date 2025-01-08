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

		response := helper.ResponseMessage("Created Order Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, ok := ctx.Get("userID")
	if !ok {
		listErr := gin.H{"errors": "User  ID Not Found in context"}
		response := helper.ResponseMessage("System Error", "User  ID Not Found", http.StatusInternalServerError, listErr)
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

	expiredAt := order.ExpiredAt

	formatResponse := transaction.FormatterTRXResponse(order.Id, expiredAt)
	response := helper.ResponseMessage("Order Created", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) GetOrdersByUserId(ctx *gin.Context) {
	userID, ok := ctx.Get("userID")
	if !ok {
		listErr := gin.H{"errors": "User ID Not Found in context"}
		response := helper.ResponseMessage("System Error", "User  ID Not Found", http.StatusInternalServerError, listErr)
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

	formatResponse := transaction.FormatterAllOrderResponses(orders)
	response := helper.ResponseMessage("Get Data Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) GetOrdersByUserIdAndOrderId(ctx *gin.Context) {
	var input transaction.OrderDetailInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Get Data Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, ok := ctx.Get("userID")
	if !ok {
		listErr := gin.H{"errors": "User  ID Not Found in context"}
		response := helper.ResponseMessage("System Error", "User  ID Not Found", http.StatusInternalServerError, listErr)
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

	if order.Id == "" {
		listErr := gin.H{"errors": "Data Not Found"}
		response := helper.ResponseMessage("Get Data Failed", "Success", http.StatusOK, listErr)
		ctx.JSON(http.StatusOK, response)
		return
	}

	formatResponse := transaction.FormatterOrderResponse(order)
	response := helper.ResponseMessage("Get Data Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) PaymentNow(ctx *gin.Context) {
	var input transaction.PaymentInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Payment Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, ok := ctx.Get("userID")
	if !ok {
		listErr := gin.H{"errors": "User  ID Not Found in context"}
		response := helper.ResponseMessage("System Error", "User  ID Not Found", http.StatusInternalServerError, listErr)
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
	response := helper.ResponseMessage("Payment Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) UpdateOrderById(ctx *gin.Context) {
	var input transaction.UpdateOrderInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Update Failed", "Bad Request", http.StatusBadRequest, listErr)
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
	response := helper.ResponseMessage("Update Success", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *trxHandler) DeleteOrderById(ctx *gin.Context) {
	var input transaction.OrderDeleteInput
	err := ctx.ShouldBind(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"errors": errors}

		response := helper.ResponseMessage("Delete Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Id == "" {
		listErr := gin.H{"errors": "Order ID cannot NULL"}
		response := helper.ResponseMessage("Delete Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteOrderById(input.Id)
	if err != nil {
		listErr := gin.H{"errors": err.Error()}
		response := helper.ResponseMessage("Delete Failed", "Failed", http.StatusInternalServerError, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ResponseMessage("Delete Success", "Success", http.StatusOK, gin.H{"message": "Delete Successfully"})
	ctx.JSON(http.StatusOK, response)
}
