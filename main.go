package main

import (
	"log"
	"slash/handler"
	"slash/helper"
	"slash/product"
	"slash/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
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

	router := gin.Default()
	api := router.Group("/api/slash")
	api.POST("/registration", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/check-email", helper.TokenAuthMiddleware(), userHandler.CheckEmailAvailable)
	api.GET("/products", helper.TokenAuthMiddleware(), prdHandler.GetAllProduct)
	api.POST("/products/name", helper.TokenAuthMiddleware(), prdHandler.FindProductByName)
	api.POST("/products/id", helper.TokenAuthMiddleware(), prdHandler.FindProductById)

	router.Run()
}
