package main

import (
	"gin-gorm-curd-rest-api/controllers"
	"gin-gorm-curd-rest-api/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()

	r := gin.Default()

	userController := new(controllers.ProductController)

	productsRoute := r.Group("/products")
	{
		productsRoute.POST("/", userController.Create)
		productsRoute.GET("/", userController.ReadAll)
		productsRoute.GET("/:id", userController.ReadOne)
		productsRoute.PUT("/:id", userController.Update)
		productsRoute.DELETE("/:id", userController.Delete)
	}

  
	r.Run() 
}