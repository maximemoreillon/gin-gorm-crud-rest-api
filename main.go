package main

import (
	"fmt"
	"strconv"
	"net/http"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
	"github.com/gin-gonic/gin"
)

type Product struct {
  gorm.Model
  Code  string
  Price int
}



func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	// Migrate the schema
  db.AutoMigrate(&Product{})

	r := gin.Default()

	r.POST("/products", func(c *gin.Context) {

		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&Product{Code: newProduct.Code, Price: newProduct.Price})

		c.JSON(http.StatusOK, gin.H{
      "status": "OK",
    })

	})


  r.GET("/products", func(c *gin.Context) {

		var products []Product

		db.Find(&products)

    c.JSON(http.StatusOK, gin.H{
      "items": products,
    })

  })

	r.GET("/products/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("Error during conversion")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var product Product
		db.First(&product, id)

    c.JSON(http.StatusOK, gin.H{
      "item": product,
    })

  })

	r.PUT("/products/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("Error during conversion")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var newProperties Product
		if err := c.ShouldBindJSON(&newProperties); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var product Product
		db.First(&product, id)
		db.Model(&product).Updates(Product{Price: newProperties.Price, Code: newProperties.Code}) // non-zero fields

		c.JSON(http.StatusOK, gin.H{"code": newProperties.Code, "price": newProperties.Price})

  })

	r.DELETE("/products/:id", func(c *gin.Context) {

		var product Product
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			fmt.Println("Error during conversion")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Delete(&product, id)

    c.JSON(http.StatusOK, gin.H{
      "id": id,
    })

  })
  

  

  // Create
  // db.Create(&Product{Code: "D42", Price: 100})

  // Read
  // var product Product
  // db.First(&product, 1) // find product with integer primary key
  // db.First(&product, "code = ?", "D42") // find product with code D42

  // // Update - update product's price to 200
  // db.Model(&product).Update("Price", 200)
  // // Update - update multiple fields
  // db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  // db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // // Delete - delete product
  // db.Delete(&product, 1)


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}