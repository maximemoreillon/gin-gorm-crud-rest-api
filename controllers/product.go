package controllers

import (
	"fmt"
	"gin-gorm-curd-rest-api/db"
	"gin-gorm-curd-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (p ProductController) Create(c *gin.Context) {

	var newProduct models.Product

	err := c.ShouldBindJSON(&newProduct); 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.GetDB().Create(&models.Product{Code: newProduct.Code, Price: newProduct.Price})

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}

func (p ProductController) ReadAll(c *gin.Context) {

	var products []models.Product

	db.GetDB().Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"items": products,
	})

}

func (p ProductController) ReadOne(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("Error during conversion")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	db.GetDB().First(&product, id)

	c.JSON(http.StatusOK, product)

}

func (p ProductController) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error during conversion")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newProperties models.Product
	if err := c.ShouldBindJSON(&newProperties); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	db.GetDB().First(&product, id)
	db.GetDB().Model(&product).Updates(models.Product{Price: newProperties.Price, Code: newProperties.Code})

	c.JSON(http.StatusOK, gin.H{"code": newProperties.Code, "price": newProperties.Price})

}
func (p ProductController) Delete(c *gin.Context) {

	var product models.Product
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("Error during conversion")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.GetDB().Delete(&product, id)

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}
