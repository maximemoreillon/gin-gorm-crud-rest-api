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

	r.POST("/products", userController.Create)
  r.GET("/products", userController.ReadAll)
	r.GET("/products/:id", userController.ReadOne)
	r.PUT("/products/:id", userController.Update)
	r.DELETE("/products/:id", userController.Delete)
  
	r.Run() 
}