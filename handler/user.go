package handler

import (
	"net/http"
	"slash/helper"
	"slash/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var input user.RegisterUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"error": errors}

		response := helper.ResponseMessage("Register Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Role == "" {
		input.Role = "helper"
	}

	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		listErr := gin.H{"error": err.Error()}
		response := helper.ResponseMessage("Register Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := helper.GenerateJWT(newUser.Id, newUser.Email, newUser.Role)
	if err != nil {
		listErr := gin.H{"error": err.Error()}
		response := helper.ResponseMessage("Login Failed", "Internal Server Error", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := user.FormatterUserResponse(newUser, token)
	response := helper.ResponseMessage("User Created", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var request user.LoginInput
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"error": errors}

		response := helper.ResponseMessage("Login Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userExist, err := h.service.Login(request)
	if err != nil {
		listErr := gin.H{"error": err.Error()}
		response := helper.ResponseMessage("Login Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := helper.GenerateJWT(userExist.Id, userExist.Email, userExist.Role)
	if err != nil {
		listErr := gin.H{"error": err.Error()}
		response := helper.ResponseMessage("Login Failed", "Internal Server Error", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatResponse := user.FormatterUserResponse(userExist, token)
	response := helper.ResponseMessage("Login Succes", "Success", http.StatusOK, formatResponse)
	ctx.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailable(ctx *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ResponseMessageValidationError(err)
		listErr := gin.H{"error": errors}

		response := helper.ResponseMessage("Check User Failed", "Bad Request", http.StatusBadRequest, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	IsEmailAvailable, err := h.service.IsEmailAvailable(input.Email)
	if err != nil {
		listErr := gin.H{"error": "Server Error"}
		response := helper.ResponseMessage("Email Is Already Use", "Access Denied", http.StatusNotAcceptable, listErr)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_available": IsEmailAvailable}
	metaMessage := "Email Is Already Use"
	if IsEmailAvailable {
		metaMessage = "Email Available"
	}
	response := helper.ResponseMessage(metaMessage, "Success", http.StatusOK, data)
	ctx.JSON(http.StatusOK, response)
}
