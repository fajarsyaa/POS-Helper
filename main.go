package main

import (
	"log"
	"slash/handler"
	"slash/helper"
	"slash/product"
	"slash/transaction"
	"slash/user"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc

	dsn := "root:@tcp(127.0.0.1:3306)/slash-helper?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	prdRepo := product.NewRepository(db)
	prdService := product.NewService(prdRepo)
	prdHandler := handler.NewProductHandler(prdService)

	trxRepo := transaction.NewRepository(db)
	trxService := transaction.NewService(trxRepo)
	trxHandler := handler.NewTransactionHandler(trxService)

	router := gin.Default()
	api := router.Group("/api/slash")
	api.POST("/registration", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailAvailable)
	api.GET("/products", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), prdHandler.GetAllProduct)
	api.POST("/products/name", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), prdHandler.FindProductByName)
	api.POST("/products/id", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), prdHandler.FindProductById)
	api.POST("/order", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), trxHandler.CreateOrder)
	api.GET("/orders", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), trxHandler.GetOrdersByUserId)
	api.POST("/order/detail", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), trxHandler.GetOrdersByUserIdAndOrderId)
	api.PUT("/order/edit", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), trxHandler.UpdateOrderById)
	api.POST("/order/payment", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), trxHandler.PaymentNow)
	api.DELETE("/order/delete", helper.TokenAuthMiddleware([]string{"helper", "pengguna"}), trxHandler.DeleteOrderById)

	router.Run()
}
